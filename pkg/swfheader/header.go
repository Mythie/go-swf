package swfheader

import (
	"fmt"
	"strconv"

	"github.com/mythie/go-swf/pkg/swftype"
)

type Signature [3]byte

type Header struct {
	Compression CompressionLevel
	Signature   Signature
	Version     uint8
	FrameRate   uint16
	FrameCount  uint16
	FileLength  uint32
	FrameSize   swftype.Rect
}

func (s *Header) FPS() float32 {
	var shift uint8 = 8

	var (
		whole uint16 = s.FrameRate >> shift
		frac  uint16 = s.FrameRate << shift
	)

	f, err := strconv.ParseFloat(fmt.Sprintf("%d.%d", whole, frac), 32)

	if err != nil {
		return 0.0
	}

	return float32(f)
}
