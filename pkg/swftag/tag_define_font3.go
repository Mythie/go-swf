package swftag

type DefineFont3Tag struct {
    *tag
}

func (t *DefineFont3Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFont3Tag) Data() []byte {
    return t.data
}

func (t *DefineFont3Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFont3Tag(id TagIdentifier, data []byte) *DefineFont3Tag {
    return &DefineFont3Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
