package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type Lossless2BitmapFormat uint8

const (
	Lossless2BitmapFormat8Bit Lossless2BitmapFormat = 3
	Lossless2BitmapFormatARGB Lossless2BitmapFormat = 5
)

type DefineBitsLossless2Tag struct {
	*tag
	bitmapData           []byte
	characterId          uint16
	bitmapWidth          uint16
	bitmapHeight         uint16
	bitmapFormat         Lossless2BitmapFormat
	bitmapColorTableSize uint8
}

func (d *DefineBitsLossless2Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":          d.tag,
		"characterId":  d.characterId,
		"bitmapWidth":  d.bitmapWidth,
		"bitmapHeight": d.bitmapHeight,
		"bitmapFormat": d.bitmapFormat,
	})
}

func NewDefineBitsLossless2Tag(id TagIdentifier, data []byte) *DefineBitsLossless2Tag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	var bitmapColorTableSize uint8

	characterId, _ := reader.ReadUInt16()
	bitmapFormat, _ := reader.ReadUInt8()
	bitmapWidth, _ := reader.ReadUInt16()
	bitmapHeight, _ := reader.ReadUInt16()

	if Lossless2BitmapFormat(bitmapFormat) == Lossless2BitmapFormat8Bit {
		bitmapColorTableSize, _ = reader.ReadUInt8()
	}

	bitmapData, _ := reader.ReadAll()

	return &DefineBitsLossless2Tag{
		tag: &tag{
			id:   id,
			data: data,
		},

		characterId:          characterId,
		bitmapFormat:         Lossless2BitmapFormat(bitmapFormat),
		bitmapWidth:          bitmapWidth,
		bitmapHeight:         bitmapHeight,
		bitmapColorTableSize: bitmapColorTableSize,
		bitmapData:           bitmapData,
	}
}
