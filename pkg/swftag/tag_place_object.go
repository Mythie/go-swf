package swftag

type PlaceObjectTag struct {
    *tag
}

func (t *PlaceObjectTag) Id() TagIdentifier {
    return t.id
}

func (t *PlaceObjectTag) Data() []byte {
    return t.data
}

func (t *PlaceObjectTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewPlaceObjectTag(id TagIdentifier, data []byte) *PlaceObjectTag {
    return &PlaceObjectTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
