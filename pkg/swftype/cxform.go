package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type CXForm struct {
	HasAddTerms   bool
	HasMultTerms  bool
	Nbits         uint8
	RedMultTerm   int64
	GreenMultTerm int64
	BlueMultTerm  int64
	RedAddTerm    int64
	GreenAddTerm  int64
	BlueAddTerm   int64
}

func ParseCXForm(reader swfreader.Reader) (*CXForm, error) {
	var (
		boolBits  uint8 = 1
		nbitsBits uint8 = 4

		redMultTerm   int64
		greenMultTerm int64
		blueMultTerm  int64
		redAddTerm    int64
		greenAddTerm  int64
		blueAddTerm   int64
	)

	hasAddTerms, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	hasMultTerms, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	nbits, err := reader.ReadUBits(nbitsBits)

	if err != nil {
		return nil, err
	}

	uNbits := uint8(nbits)

	if hasMultTerms == 1 {
		redMultTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}

		greenMultTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}

		blueMultTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}
	}

	if hasAddTerms == 1 {
		redAddTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}

		greenAddTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}

		blueAddTerm, err = reader.ReadSBits(uNbits)

		if err != nil {
			return nil, err
		}
	}

	return &CXForm{
		HasAddTerms:   hasAddTerms == 1,
		HasMultTerms:  hasMultTerms == 1,
		Nbits:         uNbits,
		RedMultTerm:   redMultTerm,
		GreenMultTerm: greenMultTerm,
		BlueMultTerm:  blueMultTerm,
		RedAddTerm:    redAddTerm,
		GreenAddTerm:  greenAddTerm,
		BlueAddTerm:   blueAddTerm,
	}, nil
}
