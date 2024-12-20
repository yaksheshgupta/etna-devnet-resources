package lib

import "github.com/ava-labs/avalanchego/utils/units"

const (
	ETNA_RPC_URL = "https://etna.avax-dev.network"
	MIN_BALANCE  = units.Avax*VALIDATORS_COUNT + 100*units.MilliAvax
	FAUCET_LINK  = "https://core.app/tools/testnet-faucet/?subnet=cdevnet&token=cdevnet"

	VALIDATOR_MANAGER_OWNER_KEY_PATH = "data/poa_validator_manager_owner_key.txt"

	VALIDATORS_COUNT = 1
	NETWORK_ID       = 76
)
