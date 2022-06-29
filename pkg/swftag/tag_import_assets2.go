package swftag

type ImportAssets2Tag struct {
    *tag
}

func (t *ImportAssets2Tag) Id() TagIdentifier {
    return t.id
}

func (t *ImportAssets2Tag) Data() []byte {
    return t.data
}

func (t *ImportAssets2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewImportAssets2Tag(id TagIdentifier, data []byte) *ImportAssets2Tag {
    return &ImportAssets2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
