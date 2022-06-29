package swftag

type VideoFrameTag struct {
    *tag
}

func (t *VideoFrameTag) Id() TagIdentifier {
    return t.id
}

func (t *VideoFrameTag) Data() []byte {
    return t.data
}

func (t *VideoFrameTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewVideoFrameTag(id TagIdentifier, data []byte) *VideoFrameTag {
    return &VideoFrameTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
