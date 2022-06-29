package swftag

type ShowFrameTag struct {
    *tag
}

func (t *ShowFrameTag) Id() TagIdentifier {
    return t.id
}

func (t *ShowFrameTag) Data() []byte {
    return t.data
}

func (t *ShowFrameTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewShowFrameTag(id TagIdentifier, data []byte) *ShowFrameTag {
    return &ShowFrameTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
