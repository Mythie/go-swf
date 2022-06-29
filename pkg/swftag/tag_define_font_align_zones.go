package swftag

type DefineFontAlignZonesTag struct {
    *tag
}

func (t *DefineFontAlignZonesTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineFontAlignZonesTag) Data() []byte {
    return t.data
}

func (t *DefineFontAlignZonesTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineFontAlignZonesTag(id TagIdentifier, data []byte) *DefineFontAlignZonesTag {
    return &DefineFontAlignZonesTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
