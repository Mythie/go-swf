package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type DefineBitsTag struct {
	*tag
	jpegData    []byte
	characterId uint16
}

func (d *DefineBitsTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":         d.tag,
		"characterId": d.characterId,
	})
}

func NewDefineBitsTag(id TagIdentifier, data []byte) *DefineBitsTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	characterId, _ := reader.ReadUInt16()
	jpegData, _ := reader.ReadAll()

	return &DefineBitsTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		characterId: characterId,
		jpegData:    jpegData,
	}
}
