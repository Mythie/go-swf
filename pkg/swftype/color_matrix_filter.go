package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type ColorMatrixFilter struct {
	Matrix [20]float32
}

func ParseColorMatrixFilter(reader swfreader.Reader) (*ColorMatrixFilter, error) {
	var (
		matrix [20]float32
		err    error
	)

	for i := range matrix {
		matrix[i], err = reader.ReadFloat32()

		if err != nil {
			return nil, err
		}
	}

	return &ColorMatrixFilter{
		Matrix: matrix,
	}, nil
}
