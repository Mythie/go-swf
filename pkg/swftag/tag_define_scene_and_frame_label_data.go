package swftag

type DefineSceneAndFrameLabelDataTag struct {
    *tag
}

func (t *DefineSceneAndFrameLabelDataTag) Id() TagIdentifier {
    return t.id
}

func (t *DefineSceneAndFrameLabelDataTag) Data() []byte {
    return t.data
}

func (t *DefineSceneAndFrameLabelDataTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewDefineSceneAndFrameLabelDataTag(id TagIdentifier, data []byte) *DefineSceneAndFrameLabelDataTag {
    return &DefineSceneAndFrameLabelDataTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
