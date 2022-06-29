package swfprimitive

import "fmt"

// const (
// 	Float16ExponentBias float64 = 15
// )

type Float16 uint16

// Float32 returns a float32 representation of the Float16.
func (f Float16) Float32() float32 {
	if f == 0 {
		return 0
	}

	sign := f & 0x8000
	exponent := (f & 0x7c00) >> 10
	mantissa := (f & 0x03ff) | 0x0400

	if sign != 0 {
		mantissa = -mantissa
	}

	k := 1 << exponent
	bias := 1 << 25

	return float32(mantissa) * float32(k) / float32(bias)
}

func (f Float16) Float64() float64 {
	return float64(f.Float32())
}

func (f Float16) String() string {
	return fmt.Sprintf("%f", f.Float32())
}
