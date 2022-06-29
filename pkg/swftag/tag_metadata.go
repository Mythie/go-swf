package swftag

type MetadataTag struct {
    *tag
}

func (t *MetadataTag) Id() TagIdentifier {
    return t.id
}

func (t *MetadataTag) Data() []byte {
    return t.data
}

func (t *MetadataTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewMetadataTag(id TagIdentifier, data []byte) *MetadataTag {
    return &MetadataTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
