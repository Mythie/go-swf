package swftag

type DefineTextTag struct {
    *tag
}

func (t *DefineTextTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineTextTag) Data() []byte {
    return t.data
}

func (t *DefineTextTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineTextTag(id TagIdentifier, data []byte) *DefineTextTag {
    return &DefineTextTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
