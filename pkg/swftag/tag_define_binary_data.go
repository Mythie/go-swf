package swftag

import (
	"bytes"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type DefineBinaryDataTag struct {
	*tag
	binaryData []byte
	reserved   uint32
}

func NewDefineBinaryDataTag(id TagIdentifier, data []byte) *DefineBinaryDataTag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	reserved, _ := reader.ReadUInt32()
	binaryData, _ := reader.ReadAll()

	return &DefineBinaryDataTag{
		tag: &tag{
			id:   id,
			data: data,
		},

		reserved:   reserved,
		binaryData: binaryData,
	}
}
