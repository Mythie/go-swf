package swftag

type DefineShapeTag struct {
    *tag
}

func (t *DefineShapeTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineShapeTag) Data() []byte {
    return t.data
}

func (t *DefineShapeTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineShapeTag(id TagIdentifier, data []byte) *DefineShapeTag {
    return &DefineShapeTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
