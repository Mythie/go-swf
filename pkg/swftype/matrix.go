package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type Matrix struct {
	NScaleBits     uint64
	NTranslateBits uint64
	TranslateX     int64
	TranslateY     int64
	NRotateBits    uint64
	RotateSkew0    float32
	RotateSkew1    float32
	ScaleX         float32
	ScaleY         float32
	HasScale       bool
	HasRotate      bool
}

func ParseMatrix(reader swfreader.Reader) (*Matrix, error) {
	var (
		boolBits  uint8 = 1
		nbitsBits uint8 = 5

		hasScale   uint64
		nScaleBits uint64
		scaleX     float32
		scaleY     float32

		hasRotate   uint64
		nRotateBits uint64
		rotateSkew0 float32
		rotateSkew1 float32

		nTranslateBits uint64
		translateX     int64
		translateY     int64
	)

	hasScale, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	if hasScale == 1 {
		nScaleBits, err = reader.ReadUBits(nbitsBits)

		if err != nil {
			return nil, err
		}

		scaleX, err = reader.ReadFBits(uint8(nScaleBits))

		if err != nil {
			return nil, err
		}

		scaleY, err = reader.ReadFBits(uint8(nScaleBits))

		if err != nil {
			return nil, err
		}
	}

	hasRotate, err = reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	if hasRotate == 1 {
		nRotateBits, err = reader.ReadUBits(nbitsBits)

		if err != nil {
			return nil, err
		}

		rotateSkew0, err = reader.ReadFBits(uint8(nRotateBits))

		if err != nil {
			return nil, err
		}

		rotateSkew1, err = reader.ReadFBits(uint8(nRotateBits))

		if err != nil {
			return nil, err
		}
	}

	nTranslateBits, err = reader.ReadUBits(nbitsBits)

	if err != nil {
		return nil, err
	}

	translateX, err = reader.ReadSBits(uint8(nTranslateBits))

	if err != nil {
		return nil, err
	}

	translateY, err = reader.ReadSBits(uint8(nTranslateBits))

	if err != nil {
		return nil, err
	}

	return &Matrix{
		HasScale:       hasScale == 1,
		NScaleBits:     nScaleBits,
		ScaleX:         scaleX,
		ScaleY:         scaleY,
		HasRotate:      hasRotate == 1,
		NRotateBits:    nRotateBits,
		RotateSkew0:    rotateSkew0,
		RotateSkew1:    rotateSkew1,
		NTranslateBits: nTranslateBits,
		TranslateX:     translateX,
		TranslateY:     translateY,
	}, nil
}
