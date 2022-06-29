package swftag

type ExportAssetsTag struct {
    *tag
}

func (t *ExportAssetsTag) Id() TagIdentifier {
    return t.id
}

func (t *ExportAssetsTag) Data() []byte {
    return t.data
}

func (t *ExportAssetsTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewExportAssetsTag(id TagIdentifier, data []byte) *ExportAssetsTag {
    return &ExportAssetsTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
