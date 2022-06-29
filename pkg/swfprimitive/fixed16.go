package swfprimitive

import "fmt"

const (
	fixed16Float32Divisor = 65536
)

type Fixed16 int32

func (f Fixed16) Float32() float32 {
	return float32(f) / fixed16Float32Divisor
}

func (f Fixed16) Float64() float64 {
	return float64(f.Float32())
}

func (f Fixed16) String() string {
	return fmt.Sprintf("%f", f.Float32())
}
