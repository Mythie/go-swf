package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
	"github.com/mythie/go-swf/pkg/swftype"
)

type DefineButtonSoundTag struct {
	*tag
	buttonSoundInfo2 *swftype.SoundInfo
	buttonSoundInfo0 *swftype.SoundInfo
	buttonSoundInfo1 *swftype.SoundInfo
	buttonSoundInfo3 *swftype.SoundInfo
	buttonSoundChar0 uint16
	buttonSoundChar1 uint16
	buttonSoundChar2 uint16
	buttonSoundChar3 uint16
	buttonId         uint16
}

func (d *DefineButtonSoundTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":              d.tag,
		"buttonId":         d.buttonId,
		"buttonSoundChar0": d.buttonSoundChar0,
		"buttonSoundInfo0": d.buttonSoundInfo0,
		"buttonSoundChar1": d.buttonSoundChar1,
		"buttonSoundInfo1": d.buttonSoundInfo1,
		"buttonSoundChar2": d.buttonSoundChar2,
		"buttonSoundInfo2": d.buttonSoundInfo2,
		"buttonSoundChar3": d.buttonSoundChar3,
		"buttonSoundInfo3": d.buttonSoundInfo3,
	})
}

func NewDefineButtonSoundTag(id TagIdentifier, data []byte) *DefineButtonSoundTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	var (
		buttonSoundInfo0 *swftype.SoundInfo
		buttonSoundInfo1 *swftype.SoundInfo
		buttonSoundInfo2 *swftype.SoundInfo
		buttonSoundInfo3 *swftype.SoundInfo
	)

	buttonId, _ := reader.ReadUInt16()
	buttonSoundChar0, _ := reader.ReadUInt16()

	if buttonSoundChar0 != 0 {
		buttonSoundInfo0, _ = swftype.ParseSoundInfo(reader)
	}

	buttonSoundChar1, _ := reader.ReadUInt16()

	if buttonSoundChar1 != 0 {
		buttonSoundInfo1, _ = swftype.ParseSoundInfo(reader)
	}

	buttonSoundChar2, _ := reader.ReadUInt16()

	if buttonSoundChar2 != 0 {
		buttonSoundInfo2, _ = swftype.ParseSoundInfo(reader)
	}

	buttonSoundChar3, _ := reader.ReadUInt16()

	if buttonSoundChar3 != 0 {
		buttonSoundInfo3, _ = swftype.ParseSoundInfo(reader)
	}

	return &DefineButtonSoundTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		buttonId:         buttonId,
		buttonSoundChar0: buttonSoundChar0,
		buttonSoundInfo0: buttonSoundInfo0,
		buttonSoundChar1: buttonSoundChar1,
		buttonSoundInfo1: buttonSoundInfo1,
		buttonSoundChar2: buttonSoundChar2,
		buttonSoundInfo2: buttonSoundInfo2,
		buttonSoundChar3: buttonSoundChar3,
		buttonSoundInfo3: buttonSoundInfo3,
	}
}
