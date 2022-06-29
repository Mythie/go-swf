package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type GlowFilter struct {
	GlowColor      *RGBA
	BlurX          float32
	BlurY          float32
	Strength       float32
	InnerGlow      uint64
	Knockout       uint64
	CompositeScore uint64
	Passes         uint64
}

func ParseGlowFilter(reader swfreader.Reader) (*GlowFilter, error) {
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

	return &GlowFilter{
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
