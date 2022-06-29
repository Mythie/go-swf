package swftag

type DefineScalingGridTag struct {
    *tag
}

func (t *DefineScalingGridTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineScalingGridTag) Data() []byte {
    return t.data
}

func (t *DefineScalingGridTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineScalingGridTag(id TagIdentifier, data []byte) *DefineScalingGridTag {
    return &DefineScalingGridTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
