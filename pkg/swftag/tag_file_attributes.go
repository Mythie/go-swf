package swftag

type FileAttributesTag struct {
    *tag
}

func (t *FileAttributesTag) Id() TagIdentifier {
    return t.id
}

func (t *FileAttributesTag) Data() []byte {
    return t.data
}

func (t *FileAttributesTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewFileAttributesTag(id TagIdentifier, data []byte) *FileAttributesTag {
    return &FileAttributesTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
