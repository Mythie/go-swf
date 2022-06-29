package swftag

type DefineShape3Tag struct {
    *tag
}

func (t *DefineShape3Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineShape3Tag) Data() []byte {
    return t.data
}

func (t *DefineShape3Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineShape3Tag(id TagIdentifier, data []byte) *DefineShape3Tag {
    return &DefineShape3Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
