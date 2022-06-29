package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type DefineBitsJPEG2Tag struct {
	*tag
	jpegData    []byte
	characterId uint16
}

func (d *DefineBitsJPEG2Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":         d.tag,
		"characterId": d.characterId,
	})
}

func NewDefineBitsJPEG2Tag(id TagIdentifier, data []byte) *DefineBitsJPEG2Tag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	characterId, _ := reader.ReadUInt16()
	jpegData, _ := reader.ReadAll()

	return &DefineBitsJPEG2Tag{
		tag: &tag{
			id:   id,
			data: data,
		},

		characterId: characterId,
		jpegData:    jpegData,
	}
}
