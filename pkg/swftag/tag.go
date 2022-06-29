package swftag

import (
	"encoding/json"
)

type TagIdentifier uint8

const (
	TagEnd                          TagIdentifier = 0
	TagShowFrame                    TagIdentifier = 1
	TagDefineShape                  TagIdentifier = 2
	TagFreeCharacter                TagIdentifier = 3
	TagPlaceObject                  TagIdentifier = 4
	TagRemoveObject                 TagIdentifier = 5
	TagDefineBits                   TagIdentifier = 6
	TagDefineButton                 TagIdentifier = 7
	TagJPEGTables                   TagIdentifier = 8
	TagSetBackgroundColor           TagIdentifier = 9
	TagDefineFont                   TagIdentifier = 10
	TagDefineText                   TagIdentifier = 11
	TagDoAction                     TagIdentifier = 12
	TagDefineFontInfo               TagIdentifier = 13
	TagDefineSound                  TagIdentifier = 14
	TagStartSound                   TagIdentifier = 15
	TagDefineButtonSound            TagIdentifier = 17
	TagSoundStreamHead              TagIdentifier = 18
	TagSoundStreamBlock             TagIdentifier = 19
	TagDefineBitsLossless           TagIdentifier = 20
	TagDefineBitsJPEG2              TagIdentifier = 21
	TagDefineShape2                 TagIdentifier = 22
	TagDefineButtonCxform           TagIdentifier = 23
	TagProtect                      TagIdentifier = 24
	TagPathsArePostScript           TagIdentifier = 25
	TagPlaceObject2                 TagIdentifier = 26
	TagRemoveObject2                TagIdentifier = 28
	TagDefineShape3                 TagIdentifier = 32
	TagDefineText2                  TagIdentifier = 33
	TagDefineButton2                TagIdentifier = 34
	TagDefineBitsJPEG3              TagIdentifier = 35
	TagDefineBitsLossless2          TagIdentifier = 36
	TagDefineSprite                 TagIdentifier = 39
	TagProductInfo                  TagIdentifier = 41
	TagFrameLabel                   TagIdentifier = 43
	TagSoundStreamHead2             TagIdentifier = 45
	TagDefineMorphShape             TagIdentifier = 46
	TagDefineFont2                  TagIdentifier = 48
	TagDefineEditText               TagIdentifier = 37
	TagExportAssets                 TagIdentifier = 56
	TagImportAssets                 TagIdentifier = 57
	TagEnableDebugger               TagIdentifier = 58
	TagDoInitAction                 TagIdentifier = 59
	TagDefineVideoStream            TagIdentifier = 60
	TagVideoFrame                   TagIdentifier = 61
	TagDefineFontInfo2              TagIdentifier = 62
	TagDebugID                      TagIdentifier = 63
	TagEnableDebugger2              TagIdentifier = 64
	TagScriptLimits                 TagIdentifier = 65
	TagSetTabIndex                  TagIdentifier = 66
	TagFileAttributes               TagIdentifier = 69
	TagPlaceObject3                 TagIdentifier = 70
	TagImportAssets2                TagIdentifier = 71
	TagDoABC                        TagIdentifier = 72
	TagDefineFontAlignZones         TagIdentifier = 73
	TagCSMTextSettings              TagIdentifier = 74
	TagDefineFont3                  TagIdentifier = 75
	TagSymbolClass                  TagIdentifier = 76
	TagMetadata                     TagIdentifier = 77
	TagDefineScalingGrid            TagIdentifier = 78
	TagDoABC2                       TagIdentifier = 82
	TagDefineShape4                 TagIdentifier = 83
	TagDefineMorphShape2            TagIdentifier = 84
	TagDefineSceneAndFrameLabelData TagIdentifier = 86
	TagDefineBinaryData             TagIdentifier = 87
	TagDefineFontName               TagIdentifier = 88
	TagDefineFont4                  TagIdentifier = 91
	TagEnableTelemetry              TagIdentifier = 93
	TagUnknown                      TagIdentifier = 255
	TagNotSupported                 TagIdentifier = 255
)

type SWFTagFactory func(id TagIdentifier, data []byte) SWFTag

type SWFTag interface {
	Id() TagIdentifier
	Length() uint32
	Data() []byte
	MarshalJSON() ([]byte, error)
}

type tag struct {
	data []byte
	id   TagIdentifier
}

func (t *tag) Id() TagIdentifier {
	return t.id
}

func (t *tag) Length() uint32 {
	return uint32(len(t.data))
}

func (t *tag) Data() []byte {
	return t.data
}

func (t *tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":     t.Id(),
		"length": t.Length(),
	})
}

func NewTag(id TagIdentifier, data []byte) SWFTag {
	switch id {
	case TagEnd:
		return NewEndTag(id, data)

	case TagShowFrame:
		return NewShowFrameTag(id, data)

	case TagDefineShape:
		return NewDefineShapeTag(id, data)

	case TagFreeCharacter:
		return NewFreeCharacterTag(id, data)

	case TagPlaceObject:
		return NewPlaceObjectTag(id, data)

	case TagRemoveObject:
		return NewRemoveObjectTag(id, data)

	case TagDefineBits:
		return NewDefineBitsTag(id, data)

	case TagDefineButton:
		return NewDefineButtonTag(id, data)

	case TagJPEGTables:
		return NewJPEGTablesTag(id, data)

	case TagSetBackgroundColor:
		return NewSetBackgroundColorTag(id, data)

	case TagDefineFont:
		return NewDefineFontTag(id, data)

	case TagDefineText:
		return NewDefineTextTag(id, data)

	case TagDoAction:
		return NewDoActionTag(id, data)

	case TagDefineFontInfo:
		return NewDefineFontInfoTag(id, data)

	case TagDefineSound:
		return NewDefineSoundTag(id, data)

	case TagStartSound:
		return NewStartSoundTag(id, data)

	case TagDefineButtonSound:
		return NewDefineButtonSoundTag(id, data)

	case TagSoundStreamHead:
		return NewSoundStreamHeadTag(id, data)

	case TagSoundStreamBlock:
		return NewSoundStreamBlockTag(id, data)

	case TagDefineBitsLossless:
		return NewDefineBitsLosslessTag(id, data)

	case TagDefineBitsJPEG2:
		return NewDefineBitsJPEG2Tag(id, data)

	case TagDefineShape2:
		return NewDefineShape2Tag(id, data)

	case TagDefineButtonCxform:
		return NewDefineButtonCxformTag(id, data)

	case TagProtect:
		return NewProtectTag(id, data)

	case TagPathsArePostScript:
		return NewPathsArePostScriptTag(id, data)

	case TagPlaceObject2:
		return NewPlaceObject2Tag(id, data)

	case TagRemoveObject2:
		return NewRemoveObject2Tag(id, data)

	case TagDefineShape3:
		return NewDefineShape3Tag(id, data)

	case TagDefineText2:
		return NewDefineText2Tag(id, data)

	case TagDefineButton2:
		return NewDefineButton2Tag(id, data)

	case TagDefineBitsJPEG3:
		return NewDefineBitsJPEG3Tag(id, data)

	case TagDefineBitsLossless2:
		return NewDefineBitsLossless2Tag(id, data)

	case TagDefineSprite:
		return NewDefineSpriteTag(id, data)

	case TagProductInfo:
		return NewProductInfoTag(id, data)

	case TagFrameLabel:
		return NewFrameLabelTag(id, data)

	case TagSoundStreamHead2:
		return NewSoundStreamHead2Tag(id, data)

	case TagDefineMorphShape:
		return NewDefineMorphShapeTag(id, data)

	case TagDefineFont2:
		return NewDefineFont2Tag(id, data)

	case TagDefineEditText:
		return NewDefineEditTextTag(id, data)

	case TagExportAssets:
		return NewExportAssetsTag(id, data)

	case TagImportAssets:
		return NewImportAssetsTag(id, data)

	case TagEnableDebugger:
		return NewEnableDebuggerTag(id, data)

	case TagDoInitAction:
		return NewDoInitActionTag(id, data)

	case TagDefineVideoStream:
		return NewDefineVideoStreamTag(id, data)

	case TagVideoFrame:
		return NewVideoFrameTag(id, data)

	case TagDefineFontInfo2:
		return NewDefineFontInfo2Tag(id, data)

	case TagDebugID:
		return NewDebugIDTag(id, data)

	case TagEnableDebugger2:
		return NewEnableDebugger2Tag(id, data)

	case TagScriptLimits:
		return NewScriptLimitsTag(id, data)

	case TagSetTabIndex:
		return NewSetTabIndexTag(id, data)

	case TagFileAttributes:
		return NewFileAttributesTag(id, data)

	case TagPlaceObject3:
		return NewPlaceObject3Tag(id, data)

	case TagImportAssets2:
		return NewImportAssets2Tag(id, data)

	case TagDoABC:
		return NewDoABCTag(id, data)

	case TagDefineFontAlignZones:
		return NewDefineFontAlignZonesTag(id, data)

	case TagCSMTextSettings:
		return NewCSMTextSettingsTag(id, data)

	case TagDefineFont3:
		return NewDefineFont3Tag(id, data)

	case TagSymbolClass:
		return NewSymbolClassTag(id, data)

	case TagMetadata:
		return NewMetadataTag(id, data)

	case TagDefineScalingGrid:
		return NewDefineScalingGridTag(id, data)

	case TagDoABC2:
		return NewDoABC2Tag(id, data)

	case TagDefineShape4:
		return NewDefineShape4Tag(id, data)

	case TagDefineMorphShape2:
		return NewDefineMorphShape2Tag(id, data)

	case TagDefineSceneAndFrameLabelData:
		return NewDefineSceneAndFrameLabelDataTag(id, data)

	case TagDefineBinaryData:
		return NewDefineBinaryDataTag(id, data)

	case TagDefineFontName:
		return NewDefineFontNameTag(id, data)

	case TagDefineFont4:
		return NewDefineFont4Tag(id, data)

	case TagEnableTelemetry:
		return NewEnableTelemetryTag(id, data)

	case TagUnknown:
		return NewUnknownTag(id, data)

	default:
		return NewNotSupportedTag(id, data)
	}
}
