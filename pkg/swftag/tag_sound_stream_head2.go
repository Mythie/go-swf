package swftag

type SoundStreamHead2Tag struct {
    *tag
}

func (t *SoundStreamHead2Tag) Id() TagIdentifier {
    return t.id
}

func (t *SoundStreamHead2Tag) Data() []byte {
    return t.data
}

func (t *SoundStreamHead2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSoundStreamHead2Tag(id TagIdentifier, data []byte) *SoundStreamHead2Tag {
    return &SoundStreamHead2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
