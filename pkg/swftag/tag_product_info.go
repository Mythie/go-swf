package swftag

type ProductInfoTag struct {
    *tag
}

func (t *ProductInfoTag) Id() TagIdentifier {
    return t.id
}

func (t *ProductInfoTag) Data() []byte {
    return t.data
}

func (t *ProductInfoTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewProductInfoTag(id TagIdentifier, data []byte) *ProductInfoTag {
    return &ProductInfoTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
