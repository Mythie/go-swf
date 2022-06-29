package swftag

import "encoding/json"

type NotSupportedTag struct {
	*tag
}

func (t *NotSupportedTag) JSON() string {
	v, _ := json.Marshal(t)

	return string(v)
}

func NewNotSupportedTag(id TagIdentifier, data []byte) *NotSupportedTag {
	return &NotSupportedTag{
		tag: &tag{
			id:   id,
			data: data,
		},
	}
}
