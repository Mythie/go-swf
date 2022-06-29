package swftag

type DefineShape4Tag struct {
    *tag
}

func (t *DefineShape4Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineShape4Tag) Data() []byte {
    return t.data
}

func (t *DefineShape4Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineShape4Tag(id TagIdentifier, data []byte) *DefineShape4Tag {
    return &DefineShape4Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
