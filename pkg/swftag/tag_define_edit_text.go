package swftag

type DefineEditTextTag struct {
    *tag
}

func (t *DefineEditTextTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineEditTextTag) Data() []byte {
    return t.data
}

func (t *DefineEditTextTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineEditTextTag(id TagIdentifier, data []byte) *DefineEditTextTag {
    return &DefineEditTextTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
