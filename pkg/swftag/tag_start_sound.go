package swftag

type StartSoundTag struct {
    *tag
}

func (t *StartSoundTag) Id() TagIdentifier {
    return t.id
}

func (t *StartSoundTag) Data() []byte {
    return t.data
}

func (t *StartSoundTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewStartSoundTag(id TagIdentifier, data []byte) *StartSoundTag {
    return &StartSoundTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
