package swftag

type ScriptLimitsTag struct {
    *tag
}

func (t *ScriptLimitsTag) Id() TagIdentifier {
    return t.id
}

func (t *ScriptLimitsTag) Data() []byte {
    return t.data
}

func (t *ScriptLimitsTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewScriptLimitsTag(id TagIdentifier, data []byte) *ScriptLimitsTag {
    return &ScriptLimitsTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
