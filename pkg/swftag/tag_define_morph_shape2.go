package swftag

type DefineMorphShape2Tag struct {
    *tag
}

func (t *DefineMorphShape2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DefineMorphShape2Tag) Data() []byte {
    return t.data
}

func (t *DefineMorphShape2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineMorphShape2Tag(id TagIdentifier, data []byte) *DefineMorphShape2Tag {
    return &DefineMorphShape2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
