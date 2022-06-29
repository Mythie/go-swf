package swftag

type PathsArePostScriptTag struct {
    *tag
}

func (t *PathsArePostScriptTag) Id() TagIdentifier {
    return t.id
}

func (t *PathsArePostScriptTag) Data() []byte {
    return t.data
}

func (t *PathsArePostScriptTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewPathsArePostScriptTag(id TagIdentifier, data []byte) *PathsArePostScriptTag {
    return &PathsArePostScriptTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
