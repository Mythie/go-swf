package swftag

type EnableDebugger2Tag struct {
    *tag
}

func (t *EnableDebugger2Tag) Id() TagIdentifier {
    return t.id
}

func (t *EnableDebugger2Tag) Data() []byte {
    return t.data
}

func (t *EnableDebugger2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewEnableDebugger2Tag(id TagIdentifier, data []byte) *EnableDebugger2Tag {
    return &EnableDebugger2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
