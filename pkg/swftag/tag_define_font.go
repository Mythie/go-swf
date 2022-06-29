package swftag

type DefineFontTag struct {
    *tag
}

func (t *DefineFontTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFontTag) Data() []byte {
    return t.data
}

func (t *DefineFontTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFontTag(id TagIdentifier, data []byte) *DefineFontTag {
    return &DefineFontTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
