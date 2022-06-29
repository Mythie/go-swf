package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type Rect struct {
	NBits uint8
	Xmin  int64
	Xmax  int64
	Ymin  int64
	Ymax  int64
}

const TwipsPerPx = 20

func (s *Rect) MinFramePx() (int64, int64) {
	return s.Xmin / TwipsPerPx, s.Ymin / TwipsPerPx
}

func (s *Rect) MaxFramePx() (int64, int64) {
	return s.Xmax / TwipsPerPx, s.Ymax / TwipsPerPx
}

func ParseRect(reader swfreader.Reader) (*Rect, error) {
	var (
		rect Rect
		err  error

		lengthBits uint8 = 5
	)

	nbits, err := reader.ReadUBits(lengthBits)

	if err != nil {
		return nil, err
	}

	rect.NBits = uint8(nbits)

	rect.Xmin, err = reader.ReadSBits(rect.NBits)

	if err != nil {
		return nil, err
	}

	rect.Xmax, err = reader.ReadSBits(rect.NBits)

	if err != nil {
		return nil, err
	}

	rect.Ymin, err = reader.ReadSBits(rect.NBits)

	if err != nil {
		return nil, err
	}

	rect.Ymax, err = reader.ReadSBits(rect.NBits)

	if err != nil {
		return nil, err
	}

	return &rect, err
}
