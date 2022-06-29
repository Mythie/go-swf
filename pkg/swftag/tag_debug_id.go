package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type DebugIDTag struct {
	*tag
	uuid []byte
}

// MarshalJSON is the custom marshaler for DebugIDTag.
func (d DebugIDTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":  d.tag,
		"uuid": d.uuid,
	})
}

func NewDebugIDTag(id TagIdentifier, data []byte) *DebugIDTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	var uuidBytes uint32 = 16

	uuid, _ := reader.ReadBytes(uuidBytes)

	return &DebugIDTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		uuid: uuid,
	}
}
