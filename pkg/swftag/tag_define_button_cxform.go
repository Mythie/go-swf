package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
	"github.com/mythie/go-swf/pkg/swftype"
)

type DefineButtonCxformTag struct {
	*tag
	cxform   *swftype.CXForm
	buttonId uint16
}

func (d *DefineButtonCxformTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":      d.tag,
		"buttonId": d.buttonId,
		"cxform":   d.cxform,
	})
}

func NewDefineButtonCxformTag(id TagIdentifier, data []byte) *DefineButtonCxformTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	buttonId, _ := reader.ReadUInt16()
	cxform, _ := swftype.ParseCXForm(reader)

	return &DefineButtonCxformTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		buttonId: buttonId,
		cxform:   cxform,
	}
}
