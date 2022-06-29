package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type SoundEnvelope struct {
	Pos44 uint32
	Left  uint16
	Right uint16
}

func ParseSoundEnvelope(reader swfreader.Reader) (*SoundEnvelope, error) {
	pos44, err := reader.ReadUInt32()

	if err != nil {
		return nil, err
	}

	left, err := reader.ReadUInt16()

	if err != nil {
		return nil, err
	}

	right, err := reader.ReadUInt16()

	if err != nil {
		return nil, err
	}

	return &SoundEnvelope{
		Pos44: pos44,
		Left:  left,
		Right: right,
	}, nil
}
