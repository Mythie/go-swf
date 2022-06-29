package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type SoundInfo struct {
	EnvelopeRecords []*SoundEnvelope
	Reserved        uint64
	InPoint         uint32
	OutPoint        uint32
	LoopCount       uint16
	HasLoops        bool
	HasInPoint      bool
	SyncStop        bool
	SyncNoMultiple  bool
	HasEnvelope     bool
	EnvPoints       uint8
	HasOutPoint     bool
}

func ParseSoundInfo(reader swfreader.Reader) (*SoundInfo, error) {
	var (
		reservedBits uint8 = 2
		boolBits     uint8 = 1

		envelopeRecords []*SoundEnvelope
		inPoint         uint32
		outPoint        uint32
		loopCount       uint16
		envPoints       uint8
	)

	reserved, err := reader.ReadUBits(reservedBits)

	if err != nil {
		return nil, err
	}

	syncStop, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	syncNoMultiple, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	hasEnvelope, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	hasLoops, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	hasOutPoint, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	hasInPoint, err := reader.ReadUBits(boolBits)

	if err != nil {
		return nil, err
	}

	if hasInPoint != 0 {
		inPoint, err = reader.ReadUInt32()

		if err != nil {
			return nil, err
		}
	}

	if hasOutPoint != 0 {
		outPoint, err = reader.ReadUInt32()

		if err != nil {
			return nil, err
		}
	}

	if hasLoops != 0 {
		loopCount, err = reader.ReadUInt16()

		if err != nil {
			return nil, err
		}
	}

	if hasEnvelope != 0 {
		envPoints, err = reader.ReadUInt8()

		if err != nil {
			return nil, err
		}
	}

	for i := uint8(0); i < envPoints; i++ {
		envelopeRecord, err := ParseSoundEnvelope(reader)

		if err != nil {
			return nil, err
		}

		envelopeRecords = append(envelopeRecords, envelopeRecord)
	}

	return &SoundInfo{
		Reserved:        reserved,
		SyncStop:        syncStop != 0,
		SyncNoMultiple:  syncNoMultiple != 0,
		HasEnvelope:     hasEnvelope != 0,
		HasLoops:        hasLoops != 0,
		HasOutPoint:     hasOutPoint != 0,
		HasInPoint:      hasInPoint != 0,
		InPoint:         inPoint,
		OutPoint:        outPoint,
		LoopCount:       loopCount,
		EnvPoints:       envPoints,
		EnvelopeRecords: envelopeRecords,
	}, nil
}
