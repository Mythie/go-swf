package swftag

type PlaceObject2Tag struct {
    *tag
}

func (t *PlaceObject2Tag) Id() TagIdentifier {
    return t.id
}

func (t *PlaceObject2Tag) Data() []byte {
    return t.data
}

func (t *PlaceObject2Tag) Length() uint32 {
    return uint32(len(t.data))
}

func NewPlaceObject2Tag(id TagIdentifier, data []byte) *PlaceObject2Tag {
    return &PlaceObject2Tag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
