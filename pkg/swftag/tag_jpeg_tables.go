package swftag

type JPEGTablesTag struct {
    *tag
}

func (t *JPEGTablesTag) Id() TagIdentifier {
    return t.id
}

func (t *JPEGTablesTag) Data() []byte {
    return t.data
}

func (t *JPEGTablesTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewJPEGTablesTag(id TagIdentifier, data []byte) *JPEGTablesTag {
    return &JPEGTablesTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
