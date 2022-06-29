package swftag

type DefineShape2Tag struct {
    *tag
}

func (t *DefineShape2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineShape2Tag) Data() []byte {
    return t.data
}

func (t *DefineShape2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineShape2Tag(id TagIdentifier, data []byte) *DefineShape2Tag {
    return &DefineShape2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
