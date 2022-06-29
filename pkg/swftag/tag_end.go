package swftag

type EndTag struct {
    *tag
}

func (t *EndTag) Id() TagIdentifier {
    return t.id
}

func (t *EndTag) Data() []byte {
    return t.data
}

func (t *EndTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewEndTag(id TagIdentifier, data []byte) *EndTag {
    return &EndTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
