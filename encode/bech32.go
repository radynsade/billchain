package encode

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Bech32Alphabet       = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
	Bech32Separator      = '1'
	Bech32PolymodTarget  = 1
	Bech32MPolymodTarget = 0x2bc830a3
)

type Bech32Encoding uint8

const (
	Bech32DefaultEncoding = iota
	Bech32MEncoding
)

var (
	ErrInvalidBech32 = errors.New("invalid bech32")
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
		Alphabet:  Bech32Alphabet,
		Separator: '1',
	}
}

// Decodes a bech32-like string into a Bech32 struct.
func DecodeBech32Like(value string, options Bech32Options) (Bech32, error) {
	var hasLower, hasUpper bool

	for i := 0; i < len(value); i++ {
		c := value[i]

		if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
	}

	if hasLower && hasUpper {
		return Bech32{}, fmt.Errorf("%w: mixed case", ErrInvalidBech32)
	}

	if hasUpper {
		value = strings.ToLower(value)
	}

	separatorPosition := strings.LastIndexByte(value, options.Separator)

	if separatorPosition == -1 {
		return Bech32{}, fmt.Errorf("%w: no separator", ErrInvalidBech32)
	}

	if separatorPosition == 0 {
		return Bech32{}, fmt.Errorf("%w: no hrp", ErrInvalidBech32)
	}

	if separatorPosition+7 > len(value) {
		return Bech32{}, fmt.Errorf("%w: no checksum", ErrInvalidBech32)
	}

	hrp := value[:separatorPosition]
	dataPart := value[separatorPosition+1:]
	var reverse [128]int

	for i := range reverse {
		reverse[i] = -1
	}
}

func DecodeBech32(value string) (Bech32, error) {
	return DecodeBech32Like(value, Bech32Options{
		Alphabet:  Bech32Alphabet,
		Separator: Bech32Separator,
	})
}

func bech32HrpExpand(hrp string) []int {
	l := len(hrp)
	out := make([]int, 0, l*2+1)

	for i := range l {
		out = append(out, int(hrp[i])>>5)
	}

	out = append(out, 0)

	for i := range l {
		out = append(out, int(hrp[i])&31)
	}

	return out
}

func bech32Polymod(values []int) int {
	chk := 1

	for _, v := range values {
		b := chk >> 25
		chk = ((chk & 0x1ffffff) << 5) ^ v

		if (b & 1) != 0 {
			chk ^= 0x3b6a57b2
		}

		if (b & 2) != 0 {
			chk ^= 0x26508e6d
		}

		if (b & 4) != 0 {
			chk ^= 0x1ea119fa
		}

		if (b & 8) != 0 {
			chk ^= 0x3d4233dd
		}

		if (b & 16) != 0 {
			chk ^= 0x2a1462b3
		}
	}

	return chk
}
