package swftag

type DefineButtonTag struct {
	*tag
}

func NewDefineButtonTag(id TagIdentifier, data []byte) *DefineButtonTag {
	return &DefineButtonTag{
		tag: &tag{
			id:   id,
			data: data,
		},
	}
}
