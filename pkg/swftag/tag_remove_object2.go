package swftag

type RemoveObject2Tag struct {
    *tag
}

func (t *RemoveObject2Tag) Id() TagIdentifier {
    return t.id
}

func (t *RemoveObject2Tag) Data() []byte {
    return t.data
}

func (t *RemoveObject2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewRemoveObject2Tag(id TagIdentifier, data []byte) *RemoveObject2Tag {
    return &RemoveObject2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
