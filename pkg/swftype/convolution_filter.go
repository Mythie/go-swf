package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type ConvolutionFilter struct {
	MatrixX       uint8
	MatrixY       uint8
	Divisor       float32
	Bias          float32
	Matrix        []float32
	DefaultColor  *RGBA
	Reserved      uint64
	Clamp         uint64
	PreserveAlpha uint64
}

func ParseConvolutionFilter(reader swfreader.Reader) (*ConvolutionFilter, error) {
	var (
		boolBits     uint8 = 1
		reservedBits uint8 = 6
	)

	matrixX, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	matrixY, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	divisor, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	bias, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	matrix := make([]float32, matrixX*matrixY)

	for i := range matrix {
		matrix[i], err = reader.ReadFloat32()

		if err != nil {
			return nil, err
		}
	}

	defaultColor, err := ParseRGBA(reader)

	if err != nil {
		return nil, err
	}

	reserved, err := reader.ReadUBits(reservedBits)

	if err != nil {
		return nil, err
	}

	clamp, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	preserveAlpha, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	return &ConvolutionFilter{
		MatrixX:       matrixX,
		MatrixY:       matrixY,
		Divisor:       divisor,
		Bias:          bias,
		Matrix:        matrix,
		DefaultColor:  defaultColor,
		Reserved:      reserved,
		Clamp:         clamp,
		PreserveAlpha: preserveAlpha,
	}, nil
}
