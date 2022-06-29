package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func ParseRGB(reader swfreader.Reader) (*RGB, error) {
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

	return &RGB{
		Red:   red,
		Green: green,
		Blue:  blue,
	}, nil
}
