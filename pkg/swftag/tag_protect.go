package swftag

type ProtectTag struct {
    *tag
}

func (t *ProtectTag) Id() TagIdentifier {
    return t.id
}

func (t *ProtectTag) Data() []byte {
    return t.data
}

func (t *ProtectTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewProtectTag(id TagIdentifier, data []byte) *ProtectTag {
    return &ProtectTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
