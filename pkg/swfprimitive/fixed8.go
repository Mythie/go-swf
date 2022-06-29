package swfprimitive

import "fmt"

const (
	fixed8Float32Divisor = 256
)

type Fixed8 int16

func (f Fixed8) Float32() float32 {
	return float32(f) / fixed8Float32Divisor
}

func (f Fixed8) Float64() float64 {
	return float64(f.Float32())
}

func (f Fixed8) String() string {
	return fmt.Sprintf("%f", f.Float32())
}
