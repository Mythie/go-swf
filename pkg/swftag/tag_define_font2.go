package swftag

type DefineFont2Tag struct {
    *tag
}

func (t *DefineFont2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFont2Tag) Data() []byte {
    return t.data
}

func (t *DefineFont2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFont2Tag(id TagIdentifier, data []byte) *DefineFont2Tag {
    return &DefineFont2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
