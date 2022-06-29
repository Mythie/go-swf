package swftag

type FrameLabelTag struct {
    *tag
}

func (t *FrameLabelTag) Id() TagIdentifier {
    return t.id
}

func (t *FrameLabelTag) Data() []byte {
    return t.data
}

func (t *FrameLabelTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewFrameLabelTag(id TagIdentifier, data []byte) *FrameLabelTag {
    return &FrameLabelTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
