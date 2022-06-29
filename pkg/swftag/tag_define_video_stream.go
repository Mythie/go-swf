package swftag

type DefineVideoStreamTag struct {
    *tag
}

func (t *DefineVideoStreamTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineVideoStreamTag) Data() []byte {
    return t.data
}

func (t *DefineVideoStreamTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineVideoStreamTag(id TagIdentifier, data []byte) *DefineVideoStreamTag {
    return &DefineVideoStreamTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
