// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"mypkg/lib"
	"os"
	"path/filepath"
	"time"

	"github.com/ava-labs/avalanche-cli/cmd/blockchaincmd"
	"github.com/ava-labs/avalanche-cli/pkg/constants"
	"github.com/ava-labs/avalanche-cli/pkg/key"
	"github.com/ava-labs/avalanche-cli/pkg/models"
	validatorManagerSDK "github.com/ava-labs/avalanche-cli/sdk/validatormanager"
	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
	goethereumcommon "github.com/ethereum/go-ethereum/common"
)

func main() {
	convertTxIDPath := filepath.Join("data", "convert_subnet_tx_id.txt")
	if txIDBytes, err := os.ReadFile(convertTxIDPath); err == nil {
		log.Fatalf("❌ Subnet was already converted in transaction: %s", string(txIDBytes))
	}

	chainIDFilePath := filepath.Join("data", "chain.txt")
	chainIDBytes, err := os.ReadFile(chainIDFilePath)
	if err != nil {
		log.Fatalf("❌ Failed to read chain ID file: %s\n", err)
	}
	chainID := ids.FromStringOrPanic(string(chainIDBytes))

	privKey, err := lib.LoadKeyFromFile(lib.VALIDATOR_MANAGER_OWNER_KEY_PATH)
	if err != nil {
		log.Fatalf("❌ Failed to load key from file: %s\n", err)
	}
	kc := secp256k1fx.NewKeychain(privKey)

	subnetIDBytes, err := os.ReadFile("data/subnet.txt")
	if err != nil {
		log.Fatalf("❌ Failed to read subnet ID file: %s\n", err)
	}
	subnetID := ids.FromStringOrPanic(string(subnetIDBytes))

	wallet, err := primary.MakeWallet(context.Background(), &primary.WalletConfig{
		URI:          lib.ETNA_RPC_URL,
		AVAXKeychain: kc,
		EthKeychain:  kc,
		SubnetIDs:    []ids.ID{subnetID},
	})
	if err != nil {
		log.Fatalf("❌ Failed to initialize wallet: %s\n", err)
	}

	softKey, err := key.NewSoft(lib.NETWORK_ID, key.WithPrivateKey(privKey))
	if err != nil {
		log.Fatalf("❌ Failed to create change owner address: %s\n", err)
	}

	changeOwnerAddress := softKey.P()[0]
	fmt.Printf("Using changeOwnerAddress: %s\n", changeOwnerAddress)

	subnetAuthKeys, err := address.ParseToIDs([]string{changeOwnerAddress})
	if err != nil {
		log.Fatalf("❌ Failed to parse subnet auth keys: %s\n", err)
	}

	validators := []models.SubnetValidator{}
	for nodeNumber := 0; nodeNumber < lib.VALIDATORS_COUNT; nodeNumber++ {
		configBytes, err := os.ReadFile(filepath.Join("data", "configs", fmt.Sprintf("config-node%d.json", nodeNumber)))
		if err != nil {
			log.Fatalf("❌ Failed to read config file: %s\n", err)
		}
		nodeConfig := lib.NodeConfig{}
		err = json.Unmarshal(configBytes, &nodeConfig)
		if err != nil {
			log.Fatalf("❌ Failed to unmarshal config: %s\n", err)
		}

		endpoint := fmt.Sprintf("http://%s:%s", nodeConfig.PublicIP, nodeConfig.HTTPPort)

		nodeID, proofOfPossession, err := getNodeInfoRetry(endpoint)
		if err != nil {
			log.Fatalf("❌ Failed to get node info: %s\n", err)
		}
		publicKey := "0x" + hex.EncodeToString(proofOfPossession.PublicKey[:])
		pop := "0x" + hex.EncodeToString(proofOfPossession.ProofOfPossession[:])

		validator := models.SubnetValidator{
			NodeID:               nodeID.String(),
			Weight:               constants.BootstrapValidatorWeight,
			Balance:              constants.BootstrapValidatorBalance,
			BLSPublicKey:         publicKey,
			BLSProofOfPossession: pop,
			ChangeOwnerAddr:      changeOwnerAddress,
		}
		validators = append(validators, validator)
	}

	avaGoBootstrapValidators, err := blockchaincmd.ConvertToAvalancheGoSubnetValidator(validators)
	if err != nil {
		log.Fatalf("❌ Failed to convert to AvalancheGo subnet validator: %s\n", err)
	}

	managerAddress := goethereumcommon.HexToAddress(validatorManagerSDK.ValidatorContractAddress)
	options := getMultisigTxOptions(subnetAuthKeys, kc)
	tx, err := wallet.P().IssueConvertSubnetTx(
		subnetID,
		chainID,
		managerAddress.Bytes(),
		avaGoBootstrapValidators,
		options...,
	)
	if err != nil {
		log.Fatalf("❌ Failed to create convert subnet tx: %s\n", err)
	}

	err = os.WriteFile(convertTxIDPath, []byte(tx.ID().String()), 0644)
	if err != nil {
		log.Fatalf("❌ Failed to save convert subnet tx ID: %s\n", err)
	}

	fmt.Printf("✅ Convert subnet tx ID: %s\n", tx.ID().String())
}

// Naively retries getting node info from the node until it succeeds
func getNodeInfoRetry(endpoint string) (nodeID ids.NodeID, proofOfPossession *signer.ProofOfPossession, err error) {
	infoClient := info.NewClient(endpoint)
	fmt.Printf("Getting node info from %s\n", endpoint)

	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		nodeID, proofOfPossession, err = infoClient.GetNodeID(ctx)
		if err == nil {
			return
		}
		time.Sleep(time.Duration(i) * time.Second)
	}
	return ids.NodeID{}, nil, fmt.Errorf("failed to get node info after 10 retries")
}

func getMultisigTxOptions(subnetAuthKeys []ids.ShortID, kc *secp256k1fx.Keychain) []common.Option {
	options := []common.Option{}
	walletAddrs := kc.Addresses().List()
	changeAddr := walletAddrs[0]
	// addrs to use for signing
	customAddrsSet := set.Set[ids.ShortID]{}
	customAddrsSet.Add(walletAddrs...)
	customAddrsSet.Add(subnetAuthKeys...)
	options = append(options, common.WithCustomAddresses(customAddrsSet))
	// set change to go to wallet addr (instead of any other subnet auth key)
	changeOwner := &secp256k1fx.OutputOwners{
		Threshold: 1,
		Addrs:     []ids.ShortID{changeAddr},
	}
	options = append(options, common.WithChangeOwner(changeOwner))
	return options
}
