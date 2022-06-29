package swftype

import (
	"github.com/mythie/go-swf/pkg/swfprimitive"
	"github.com/mythie/go-swf/pkg/swfreader"
)

type GradientGlowFilter struct {
	NumColors     uint8
	Colors        []*RGBA
	GradientRatio []uint8
	BlurX         float32
	BlurY         float32
	Angle         float32
	Distance      float32
	Strength      swfprimitive.Fixed8
	InnerShadow
}

func ParseGradientGlowFilter(reader swfreader.Reader) (*GradientGlowFilter, error) {
	var (
		passesBits uint8 = 5
		flagBits   uint8 = 1
	)

	glowColor, err := ParseRGBA(reader)

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

	strength, err := reader.ReadFloat32()

	if err != nil {
		return nil, err
	}

	innerGlow, err := reader.ReadUBits(flagBits)

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

	return &GradientGlowFilter{
		GlowColor:      glowColor,
		BlurX:          blurX,
		BlurY:          blurY,
		Strength:       strength,
		InnerGlow:      innerGlow,
		Knockout:       knockout,
		CompositeScore: compositeScore,
		Passes:         passes,
	}, nil
}
