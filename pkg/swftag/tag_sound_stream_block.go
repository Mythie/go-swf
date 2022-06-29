package swftag

type SoundStreamBlockTag struct {
    *tag
}

func (t *SoundStreamBlockTag) Id() TagIdentifier {
    return t.id
}

func (t *SoundStreamBlockTag) Data() []byte {
    return t.data
}

func (t *SoundStreamBlockTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSoundStreamBlockTag(id TagIdentifier, data []byte) *SoundStreamBlockTag {
    return &SoundStreamBlockTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
