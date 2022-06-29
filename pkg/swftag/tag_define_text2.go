package swftag

type DefineText2Tag struct {
    *tag
}

func (t *DefineText2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineText2Tag) Data() []byte {
    return t.data
}

func (t *DefineText2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineText2Tag(id TagIdentifier, data []byte) *DefineText2Tag {
    return &DefineText2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
