package swftag

type DefineFontInfo2Tag struct {
    *tag
}

func (t *DefineFontInfo2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFontInfo2Tag) Data() []byte {
    return t.data
}

func (t *DefineFontInfo2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFontInfo2Tag(id TagIdentifier, data []byte) *DefineFontInfo2Tag {
    return &DefineFontInfo2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
