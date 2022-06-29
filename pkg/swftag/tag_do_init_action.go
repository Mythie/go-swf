package swftag

type DoInitActionTag struct {
    *tag
}

func (t *DoInitActionTag) Id() TagIdentifier {
    return t.id
}

func (t *DoInitActionTag) Data() []byte {
    return t.data
}

func (t *DoInitActionTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDoInitActionTag(id TagIdentifier, data []byte) *DoInitActionTag {
    return &DoInitActionTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
