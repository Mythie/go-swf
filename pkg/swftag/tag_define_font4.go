package swftag

type DefineFont4Tag struct {
    *tag
}

func (t *DefineFont4Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFont4Tag) Data() []byte {
    return t.data
}

func (t *DefineFont4Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFont4Tag(id TagIdentifier, data []byte) *DefineFont4Tag {
    return &DefineFont4Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
