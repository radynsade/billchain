package main

import "github.com/radynsade/billchain/encode"

const (
	Bech32BtcMainnetHrp = "bc"
	Bech32BtcTestnetHrp = "tb"
	Bech32BtcRegtestHrp = "bcrt"
	Bech32LtcMainnetHrp = "ltc"
	Bech32LtcTestnetHrp = "tltc"
)

// TODO: Add more base58 versions, ask AI for more.
const (
	Base58BtcMainnetVersion  = 0x00
	Base58BtcMainnet2Version = 0x05
	Base58BtcTestnetVersion  = 0x6F
	Base58TronMainnetVersion = 0x41
	Base58TronTestnetVersion = 0x49
)

func main() {
	encode.DecodeBech32("", encode.Bech32BitcoinOptions())
}
