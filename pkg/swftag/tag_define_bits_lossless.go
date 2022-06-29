package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type LosslessBitmapFormat uint8

const (
	LosslessBitmapFormat8Bit  LosslessBitmapFormat = 3
	LosslessBitmapFormat15Bit LosslessBitmapFormat = 4
	LosslessBitmapFormat24Bit LosslessBitmapFormat = 5
)

type DefineBitsLosslessTag struct {
	*tag
	bitmapData           []byte
	characterId          uint16
	bitmapWidth          uint16
	bitmapHeight         uint16
	bitmapFormat         LosslessBitmapFormat
	bitmapColorTableSize uint8
}

func (d *DefineBitsLosslessTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":          d.tag,
		"characterId":  d.characterId,
		"bitmapWidth":  d.bitmapWidth,
		"bitmapHeight": d.bitmapHeight,
		"bitmapFormat": d.bitmapFormat,
	})
}

func NewDefineBitsLosslessTag(id TagIdentifier, data []byte) *DefineBitsLosslessTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	var bitmapColorTableSize uint8

	characterId, _ := reader.ReadUInt16()
	bitmapFormat, _ := reader.ReadUInt8()
	bitmapWidth, _ := reader.ReadUInt16()
	bitmapHeight, _ := reader.ReadUInt16()

	if LosslessBitmapFormat(bitmapFormat) == LosslessBitmapFormat8Bit {
		bitmapColorTableSize, _ = reader.ReadUInt8()
	}

	bitmapData, _ := reader.ReadAll()

	return &DefineBitsLosslessTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		characterId:          characterId,
		bitmapFormat:         LosslessBitmapFormat(bitmapFormat),
		bitmapWidth:          bitmapWidth,
		bitmapHeight:         bitmapHeight,
		bitmapColorTableSize: bitmapColorTableSize,
		bitmapData:           bitmapData,
	}
}
