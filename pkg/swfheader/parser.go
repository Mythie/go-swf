package swfheader

import (
	"fmt"

	"github.com/mythie/go-swf/pkg/swfreader"
	"github.com/mythie/go-swf/pkg/swftype"
)

func ParseHeader(reader swfreader.Reader) (*Header, error) {
	var (
		header Header
		err    error
	)

	fmt.Println("Parsing header")

	fmt.Println("start byte:", reader.PositionBytes())
	compression, err := reader.ReadUInt8()
	fmt.Println("post-compression byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	header.Compression = CompressionLevel(compression)

	header.Signature[0] = byte(header.Compression)

	bytes, err := reader.ReadBytes(2)
	fmt.Println("post-signature byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	header.Signature[1] = bytes[0]
	header.Signature[2] = bytes[1]

	header.Version, err = reader.ReadUInt8()
	fmt.Println("post-version byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	header.FileLength, err = reader.ReadUInt32()
	fmt.Println("post-filelength byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	if header.Compression == CompressionZlib {
		err := reader.UseZlibReader()

		if err != nil {
			return nil, err
		}
	}

	if header.Compression == CompressionLzma {
		err := reader.UseLzmaReader()

		if err != nil {
			return nil, err
		}
	}
	fmt.Println("post-readerswap byte:", reader.PositionBytes())

	rect, err := swftype.ParseRect(reader)
	fmt.Println("post-rect byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	header.FrameSize = *rect

	header.FrameRate, err = reader.ReadUInt16()
	fmt.Println("post-framerate byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	header.FrameCount, err = reader.ReadUInt16()
	fmt.Println("post-framecount byte:", reader.PositionBytes())

	if err != nil {
		return nil, err
	}

	return &header, err
}
