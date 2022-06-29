package swftag

type SoundStreamHeadTag struct {
    *tag
}

func (t *SoundStreamHeadTag) Id() TagIdentifier {
    return t.id
}

func (t *SoundStreamHeadTag) Data() []byte {
    return t.data
}

func (t *SoundStreamHeadTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSoundStreamHeadTag(id TagIdentifier, data []byte) *SoundStreamHeadTag {
    return &SoundStreamHeadTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
