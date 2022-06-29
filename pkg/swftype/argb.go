package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type ARGB struct {
	Alpha uint8
	Red   uint8
	Green uint8
	Blue  uint8
}

func ParseARGB(reader swfreader.Reader) (*ARGB, error) {
	alpha, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

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

	return &ARGB{
		Alpha: alpha,
		Red:   red,
		Green: green,
		Blue:  blue,
	}, nil
}
