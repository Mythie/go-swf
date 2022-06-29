package swftag

type DoABCTag struct {
    *tag
}

func (t *DoABCTag) Id() TagIdentifier {
    return t.id
}

func (t *DoABCTag) Data() []byte {
    return t.data
}

func (t *DoABCTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDoABCTag(id TagIdentifier, data []byte) *DoABCTag {
    return &DoABCTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
