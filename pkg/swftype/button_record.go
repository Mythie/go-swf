package swftype

type ButtonRecord struct {
	ButtonReserved      int64
	ButtonHasBlendMode  bool
	ButtonHasFilterList bool
	ButtonStateHitTest  bool
	ButtonStateDown     bool
	ButtonStateOver     bool
	ButtonStateUp       bool
	CharacterId         uint16
	PlaceDepth          uint16
	PlaceMatrix         *Matrix
	ColorTransform      *CXFormWithAlpha
	FilterList          *FilterList
	BlendMode           uint8
}
