package swftag

type DefineSoundTag struct {
    *tag
}

func (t *DefineSoundTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineSoundTag) Data() []byte {
    return t.data
}

func (t *DefineSoundTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineSoundTag(id TagIdentifier, data []byte) *DefineSoundTag {
    return &DefineSoundTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
