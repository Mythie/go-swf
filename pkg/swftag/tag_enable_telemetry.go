package swftag

type EnableTelemetryTag struct {
    *tag
}

func (t *EnableTelemetryTag) Id() TagIdentifier {
    return t.id
}

func (t *EnableTelemetryTag) Data() []byte {
    return t.data
}

func (t *EnableTelemetryTag) Length() uint32 {
    return uint32(len(t.data))
}

func NewEnableTelemetryTag(id TagIdentifier, data []byte) *EnableTelemetryTag {
    return &EnableTelemetryTag{
        tag: &tag{
            id:   id,
            data: data,
        },
    }
}
