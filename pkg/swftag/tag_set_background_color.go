package swftag

type SetBackgroundColorTag struct {
    *tag
}

func (t *SetBackgroundColorTag) Id() TagIdentifier {
    return t.id
}

func (t *SetBackgroundColorTag) Data() []byte {
    return t.data
}

func (t *SetBackgroundColorTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSetBackgroundColorTag(id TagIdentifier, data []byte) *SetBackgroundColorTag {
    return &SetBackgroundColorTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
