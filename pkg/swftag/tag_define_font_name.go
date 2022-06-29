package swftag

type DefineFontNameTag struct {
    *tag
}

func (t *DefineFontNameTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFontNameTag) Data() []byte {
    return t.data
}

func (t *DefineFontNameTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFontNameTag(id TagIdentifier, data []byte) *DefineFontNameTag {
    return &DefineFontNameTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
