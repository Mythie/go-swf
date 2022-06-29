package swftag

type ImportAssetsTag struct {
    *tag
}

func (t *ImportAssetsTag) Id() TagIdentifier {
    return t.id
}

func (t *ImportAssetsTag) Data() []byte {
    return t.data
}

func (t *ImportAssetsTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewImportAssetsTag(id TagIdentifier, data []byte) *ImportAssetsTag {
    return &ImportAssetsTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
