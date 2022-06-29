package swftag

type DefineSpriteTag struct {
    *tag
}

func (t *DefineSpriteTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineSpriteTag) Data() []byte {
    return t.data
}

func (t *DefineSpriteTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineSpriteTag(id TagIdentifier, data []byte) *DefineSpriteTag {
    return &DefineSpriteTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
