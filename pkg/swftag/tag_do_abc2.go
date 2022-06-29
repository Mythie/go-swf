package swftag

type DoABC2Tag struct {
    *tag
}

func (t *DoABC2Tag) Id() TagIdentifier {
    return t.id
}

func (t *DoABC2Tag) Data() []byte {
    return t.data
}

func (t *DoABC2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDoABC2Tag(id TagIdentifier, data []byte) *DoABC2Tag {
    return &DoABC2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
