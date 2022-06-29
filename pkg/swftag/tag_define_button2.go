package swftag

type DefineButton2Tag struct {
    *tag
}

func (t *DefineButton2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineButton2Tag) Data() []byte {
    return t.data
}

func (t *DefineButton2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineButton2Tag(id TagIdentifier, data []byte) *DefineButton2Tag {
    return &DefineButton2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
