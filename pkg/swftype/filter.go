package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type FilterType uint8

const (
	DropShadowFilterType    FilterType = 0
	BlurFilterType          FilterType = 1
	GlowFilterType          FilterType = 2
	BevelFilterType         FilterType = 3
	GradientGlowFilterType  FilterType = 4
	ConvolutionFilterType   FilterType = 5
	ColorMatrixFilterType   FilterType = 6
	GradientBevelFilterType FilterType = 7
)

type Filter struct {
	FilterId            FilterType
	DropShadowFilter    *DropShadowFilter
	BlurFilter          *BlurFilter
	GlowFilter          *GlowFilter
	BevelFilter         *BevelFilter
	GradientGlowFilter  *GradientGlowFilter
	ConvolutionFilter   *ConvolutionFilter
	ColorMatrixFilter   *ColorMatrixFilter
	GradientBevelFilter *GradientBevelFilter
}

func ParseFilter(reader swfreader.Reader) (*Filter, error) {
	var (
		dropShadowFilter    *DropShadowFilter
		blurFilter          *BlurFilter
		glowFilter          *GlowFilter
		bevelFilter         *BevelFilter
		gradientGlowFilter  *GradientGlowFilter
		convolutionFilter   *ConvolutionFilter
		colorMatrixFilter   *ColorMatrixFilter
		gradientBevelFilter *GradientBevelFilter
	)

	filterId, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	switch FilterType(filterId) {
	case DropShadowFilterType:
		dropShadowFilter, err = ParseDropShadowFilter(reader)

	case BlurFilterType:
		blurFilter, err = ParseBlurFilter(reader)

	case GlowFilterType:
		glowFilter, err = ParseGlowFilter(reader)

	case BevelFilterType:
		bevelFilter, err = ParseBevelFilter(reader)

	case GradientGlowFilterType:
		gradientGlowFilter, err = ParseGradientGlowFilter(reader)

	case ConvolutionFilterType:
		convolutionFilter, err = ParseConvolutionFilter(reader)

	case ColorMatrixFilterType:
		colorMatrixFilter, err = ParseColorMatrixFilter(reader)

	case GradientBevelFilterType:
		gradientBevelFilter, err = ParseGradientBevelFilter(reader)
	}

	if err != nil {
		return nil, err
	}

	return &Filter{
		FilterId:            FilterType(filterId),
		DropShadowFilter:    dropShadowFilter,
		BlurFilter:          blurFilter,
		GlowFilter:          glowFilter,
		BevelFilter:         bevelFilter,
		GradientGlowFilter:  gradientGlowFilter,
		ConvolutionFilter:   convolutionFilter,
		ColorMatrixFilter:   colorMatrixFilter,
		GradientBevelFilter: gradientBevelFilter,
	}, nil
}
