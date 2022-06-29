package swftag

type FreeCharacterTag struct {
    *tag
}

func (t *FreeCharacterTag) Id() TagIdentifier {
    return t.id
}

func (t *FreeCharacterTag) Data() []byte {
    return t.data
}

func (t *FreeCharacterTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewFreeCharacterTag(id TagIdentifier, data []byte) *FreeCharacterTag {
    return &FreeCharacterTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
