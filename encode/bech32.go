package encode

const (
	Bech32BitcoinAlphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
	Bech32PolymodTarget   = 1
	Bech32MPolymodTarget  = 0x2bc830a3
)

type Bech32Encoding uint8

const (
	Bech32DefaultEncoding = iota
	Bech32MEncoding
)

type Bech32 struct {
	hrp      string
	data     []byte
	encoding Bech32Encoding
}

func (b Bech32) Hrp() string {
	return b.hrp
}

func (b Bech32) Data() []byte {
	return b.data
}

func (b Bech32) Encoding() Bech32Encoding {
	return b.encoding
}

type Bech32Options struct {
	Alphabet  string
	Separator byte
}

func Bech32BitcoinOptions() Bech32Options {
	return Bech32Options{
		Alphabet:  Bech32BitcoinAlphabet,
		Separator: '1',
	}
}

func DecodeBech32(value string, options Bech32Options) Bech32 {
	return Bech32{}
}
