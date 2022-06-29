package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type BlurFilter struct {
	BlurX    float32
	BlurY    float32
	Passes   uint64
	Reserved uint64
}

func ParseBlurFilter(reader swfreader.Reader) (*BlurFilter, error) {
	var (
		passesBits   uint8 = 5
		reservedBits uint8 = 3
	)

	blurX, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	blurY, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	passes, err := reader.ReadUBits(passesBits)

	if err != nil {
		return nil, err
	}

	reserved, err := reader.ReadUBits(reservedBits)

	if err != nil {
		return nil, err
	}

	return &BlurFilter{
		BlurX:    blurX,
		BlurY:    blurY,
		Passes:   passes,
		Reserved: reserved,
	}, nil
}
