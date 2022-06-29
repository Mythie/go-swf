package swftype

import (
	"github.com/mythie/go-swf/pkg/swfprimitive"
	"github.com/mythie/go-swf/pkg/swfreader"
)

type DropShadowFilter struct {
	DropShadowColor *RGBA
	BlurX           float32
	BlurY           float32
	Angle           float32
	Distance        float32
	Strength        swfprimitive.Fixed8
	InnerShadow     uint64
	Knockout        uint64
	CompositeScore  uint64
	Passes          uint64
}

func ParseDropShadowFilter(reader swfreader.Reader) (*DropShadowFilter, error) {
	var (
		passesBits uint8 = 5
		flagBits   uint8 = 1
	)

	dropShadowColor, err := ParseRGBA(reader)

	if err != nil {
		return nil, err
	}

	blurX, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	blurY, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	angle, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	distance, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	strength, err := reader.ReadFixed8()

	if err != nil {
		return nil, err
	}

	innerShadow, err := reader.ReadUBits(flagBits)

	if err != nil {
		return nil, err
	}

	knockout, err := reader.ReadUBits(flagBits)

	if err != nil {
		return nil, err
	}

	compositeScore, err := reader.ReadUBits(flagBits)

	if err != nil {
		return nil, err
	}

	passes, err := reader.ReadUBits(passesBits)

	if err != nil {
		return nil, err
	}

	return &DropShadowFilter{
		DropShadowColor: dropShadowColor,
		BlurX:           blurX,
		BlurY:           blurY,
		Angle:           angle,
		Distance:        distance,
		Strength:        strength,
		InnerShadow:     innerShadow,
		Knockout:        knockout,
		CompositeScore:  compositeScore,
		Passes:          passes,
	}, nil
}
