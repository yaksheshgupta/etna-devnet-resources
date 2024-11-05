package lib

import "github.com/ava-labs/avalanchego/utils/units"

const (
	ETNA_RPC_URL = "https://etna.avax-dev.network"
	MIN_BALANCE  = units.Avax*VALIDATORS_COUNT + 100*units.MilliAvax
	FAUCET_LINK  = "https://core.app/tools/testnet-faucet/?subnet=cdevnet&token=cdevnet"

	VALIDATOR_MANAGER_OWNER_KEY_PATH = "data/poa_validator_manager_owner_key.txt"
	TELEPORTER_DEPLOYER_KEY_PATH     = "data/teleporter_deployer_key.txt"

	VALIDATORS_COUNT = 5
	NETWORK_ID       = 76
)