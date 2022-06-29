package swftag

type SetTabIndexTag struct {
    *tag
}

func (t *SetTabIndexTag) Id() TagIdentifier {
    return t.id
}

func (t *SetTabIndexTag) Data() []byte {
    return t.data
}

func (t *SetTabIndexTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSetTabIndexTag(id TagIdentifier, data []byte) *SetTabIndexTag {
    return &SetTabIndexTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
