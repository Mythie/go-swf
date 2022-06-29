package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type RGBA struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func ParseRGBA(reader swfreader.Reader) (*RGBA, error) {
	red, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	green, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	blue, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	alpha, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	return &RGBA{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}, nil
}
