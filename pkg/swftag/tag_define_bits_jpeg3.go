package swftag

import (
	"bytes"
	"encoding/json"

	"github.com/mythie/go-swf/pkg/swfreader"
)

type DefineBitsJPEG3Tag struct {
	*tag
	jpegData        []byte
	bitmapAlphaData []byte
	alphaDataOffset uint32
	characterId     uint16
}

func (d *DefineBitsJPEG3Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"tag":             d.tag,
		"characterId":     d.characterId,
		"alphaDataOffset": d.alphaDataOffset,
	})
}

func NewDefineBitsJPEG3Tag(id TagIdentifier, data []byte) *DefineBitsJPEG3Tag {
	reader, _ := swfreader.NewReader(bytes.NewReader(data))

	// The uint16 and uint32 reads will consume 6 bytes,
	// therefore we need to subtract 6 from the length of data
	// when reading the jpegData
	var characterIdAndAlphaDataOffset uint32 = 6

	characterId, _ := reader.ReadUInt16()
	alphaDataOffset, _ := reader.ReadUInt32()

	jpegDataLen := uint32(len(data)) - characterIdAndAlphaDataOffset - alphaDataOffset

	jpegData, _ := reader.ReadBytes(jpegDataLen)

	bitmapAlphaData, _ := reader.ReadAll()

	return &DefineBitsJPEG3Tag{
		tag: &tag{
			id:   id,
			data: data,
		},

		characterId:     characterId,
		alphaDataOffset: alphaDataOffset,
		jpegData:        jpegData,
		bitmapAlphaData: bitmapAlphaData,
	}
}
