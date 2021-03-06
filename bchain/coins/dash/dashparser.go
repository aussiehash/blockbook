package dash

import (
	"blockbook/bchain/coins/btc"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

const (
	MainnetMagic wire.BitcoinNet = 0xbd6b0cbf
	TestnetMagic wire.BitcoinNet = 0xffcae2ce
	RegtestMagic wire.BitcoinNet = 0xdcb7c1fc
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
	RegtestParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	// Address encoding magics
	MainNetParams.PubKeyHashAddrID = 76 // base58 prefix: X
	MainNetParams.ScriptHashAddrID = 16 // base58 prefix: 7

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic

	// Address encoding magics
	TestNetParams.PubKeyHashAddrID = 140 // base58 prefix: y
	TestNetParams.ScriptHashAddrID = 19  // base58 prefix: 8 or 9

	RegtestParams = chaincfg.RegressionNetParams
	RegtestParams.Net = RegtestMagic

	// Address encoding magics
	RegtestParams.PubKeyHashAddrID = 140 // base58 prefix: y
	RegtestParams.ScriptHashAddrID = 19  // base58 prefix: 8 or 9

	err := chaincfg.Register(&MainNetParams)
	if err == nil {
		err = chaincfg.Register(&TestNetParams)
	}
	if err == nil {
		err = chaincfg.Register(&RegtestParams)
	}
	if err != nil {
		panic(err)
	}
}

// DashParser handle
type DashParser struct {
	*btc.BitcoinParser
}

// NewDashParser returns new DashParser instance
func NewDashParser(params *chaincfg.Params, c *btc.Configuration) *DashParser {
	return &DashParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Dash network,
// the regression test Dash network, the test Dash network and
// the simulation test Dash network, in this order
func GetChainParams(chain string) *chaincfg.Params {
	switch chain {
	case "test":
		return &TestNetParams
	case "regtest":
		return &RegtestParams
	default:
		return &MainNetParams
	}
}
