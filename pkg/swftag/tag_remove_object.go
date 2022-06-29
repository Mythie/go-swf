package swftag

type RemoveObjectTag struct {
    *tag
}

func (t *RemoveObjectTag) Id() TagIdentifier {
    return t.id
}

func (t *RemoveObjectTag) Data() []byte {
    return t.data
}

func (t *RemoveObjectTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewRemoveObjectTag(id TagIdentifier, data []byte) *RemoveObjectTag {
    return &RemoveObjectTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
