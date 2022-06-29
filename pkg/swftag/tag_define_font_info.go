package swftag

type DefineFontInfoTag struct {
    *tag
}

func (t *DefineFontInfoTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFontInfoTag) Data() []byte {
    return t.data
}

func (t *DefineFontInfoTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFontInfoTag(id TagIdentifier, data []byte) *DefineFontInfoTag {
    return &DefineFontInfoTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
