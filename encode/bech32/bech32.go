package bech32

import "errors"

var (
	ErrInvalidBech32 = errors.New("invalid bech32")
	ErrNoSeparator   = errors.New("no separator")
	ErrNoHrp         = errors.New("no hrp")
	ErrNoChecksum    = errors.New("no checksum")
)

type Bech32 struct {
	hrp  string
	data []byte
}

type Encoding struct {
	alphabet  string
	separator byte
}

func NewEncoding(alphabet string, separator byte) *Encoding {
	return &Encoding{alphabet, separator}
}

func (e *Encoding) Encode() {

}

func (e *Encoding) Decode() {

}

func hrpExpand(hrp string) []int {
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

func polymod(values []int) int {
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

var (
	StdEncoding = NewEncoding("qpzry9x8gf2tvdw0s3jn54khce6mua7l", '1')
)
