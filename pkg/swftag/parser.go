package swftag

import (
	"github.com/mythie/go-swf/pkg/swfreader"
)

// parseTag reads a tag from the reader.
func parseTag(reader swfreader.Reader) (SWFTag, error) {
	var (
		err error

		idShift              uint8  = 6
		extendedLengthMarker uint16 = 0x3f
	)

	codeAndLength, err := reader.ReadUInt16()

	if err != nil {
		return nil, err
	}

	id := TagIdentifier(codeAndLength >> idShift)

	length := int32(codeAndLength & extendedLengthMarker)

	if length == int32(extendedLengthMarker) {
		length, err = reader.ReadInt32()

		if err != nil {
			return nil, err
		}
	}

	bytes, err := reader.ReadBytes(uint32(length))

	if err != nil {
		return nil, err
	}

	tag := NewTag(id, bytes)

	return tag, err
}

// ParseTags reads all the tags from the reader.
func ParseTags(reader swfreader.Reader) (*[]SWFTag, error) {
	var (
		tags []SWFTag
		err  error
	)

	for {
		tag, err := parseTag(reader)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)

		if tag.Id() == TagEnd {
			break
		}
	}

	return &tags, err
}
