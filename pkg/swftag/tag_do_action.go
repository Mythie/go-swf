package swftag

type DoActionTag struct {
    *tag
}

func (t *DoActionTag) Id() TagIdentifier {
    return t.id
}

func (t *DoActionTag) Data() []byte {
    return t.data
}

func (t *DoActionTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDoActionTag(id TagIdentifier, data []byte) *DoActionTag {
    return &DoActionTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
