package swftag

type SymbolClassTag struct {
    *tag
}

func (t *SymbolClassTag) Id() TagIdentifier {
    return t.id
}

func (t *SymbolClassTag) Data() []byte {
    return t.data
}

func (t *SymbolClassTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewSymbolClassTag(id TagIdentifier, data []byte) *SymbolClassTag {
    return &SymbolClassTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
