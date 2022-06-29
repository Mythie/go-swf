package swftag

type DefineMorphShapeTag struct {
    *tag
}

func (t *DefineMorphShapeTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineMorphShapeTag) Data() []byte {
    return t.data
}

func (t *DefineMorphShapeTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineMorphShapeTag(id TagIdentifier, data []byte) *DefineMorphShapeTag {
    return &DefineMorphShapeTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
