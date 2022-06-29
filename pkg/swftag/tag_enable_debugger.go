package swftag

type EnableDebuggerTag struct {
    *tag
}

func (t *EnableDebuggerTag) Id() TagIdentifier {
    return t.id
}

func (t *EnableDebuggerTag) Data() []byte {
    return t.data
}

func (t *EnableDebuggerTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewEnableDebuggerTag(id TagIdentifier, data []byte) *EnableDebuggerTag {
    return &EnableDebuggerTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
