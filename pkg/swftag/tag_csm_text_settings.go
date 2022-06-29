package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type GridFit uint8
type FlashType uint8

const (
	GridFitNone      GridFit = 0
	GridFitPixels    GridFit = 1
	GridFitSubPixels GridFit = 2
)

const (
	FlashTypeNormal   FlashType = 0
	FlashTypeAdvanced FlashType = 1
)

type CSMTextSettingsTag struct {
	*tag
	reserved     uint64
	textId       uint16
	thickness    float32
	sharpness    float32
	gridFit      GridFit
	useFlashType FlashType
	reserved2    uint8
}

func (t *CSMTextSettingsTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":          t.tag,
		"textId":       t.textId,
		"useFlashType": t.useFlashType,
		"gridFit":      t.gridFit,
		"reserved":     t.reserved,
		"thickness":    t.thickness,
		"sharpness":    t.sharpness,
		"reserved2":    t.reserved2,
	})
}

func NewCSMTextSettingsTag(id TagIdentifier, data []byte) *CSMTextSettingsTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	var (
		flashTypeBits uint8 = 2
		gridFitBits   uint8 = 3
		reservedBits  uint8 = 3
	)

	textId, _ := reader.ReadUInt16()
	useFlashType, _ := reader.ReadUBits(flashTypeBits)
	gridFit, _ := reader.ReadUBits(gridFitBits)
	reserved, _ := reader.ReadUBits(reservedBits)
	thickness, _ := reader.ReadFloat32()
	sharpness, _ := reader.ReadFloat32()
	reserved2, _ := reader.ReadUInt8()

	t := &CSMTextSettingsTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		textId:       textId,
		useFlashType: FlashType(useFlashType),
		gridFit:      GridFit(gridFit),
		reserved:     reserved,
		thickness:    thickness,
		sharpness:    sharpness,
		reserved2:    reserved2,
	}

	return t
}
