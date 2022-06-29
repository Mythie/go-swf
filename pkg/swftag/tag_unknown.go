package swftag

type UnknownTag struct {
	*tag
}

func (t *UnknownTag) MarshalJSON() ([]byte, error) {
	return []byte(`"unknown"`), nil
}

func NewUnknownTag(id TagIdentifier, data []byte) *UnknownTag {
	return &UnknownTag{
		tag: &tag{
			id:   id,
			data: data,
		},
	}
}
