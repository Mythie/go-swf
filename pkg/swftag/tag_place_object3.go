package swftag

type PlaceObject3Tag struct {
    *tag
}

func (t *PlaceObject3Tag) Id() TagIdentifier {
    return t.id
}

func (t *PlaceObject3Tag) Data() []byte {
    return t.data
}

func (t *PlaceObject3Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewPlaceObject3Tag(id TagIdentifier, data []byte) *PlaceObject3Tag {
    return &PlaceObject3Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
