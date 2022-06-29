package swfreader

import (
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/icza/bitio"
	"github.com/mythie/go-swf/pkg/swfprimitive"
	"github.com/ulikunitz/xz/lzma"
)

const (
	SwfSignatureOffset = 8

	byteBits = 8
)

type Reader interface {
	PositionBytes() uint64
	PositionBits() uint64

	UseZlibReader() error
	UseLzmaReader() error
	// readLE(dst interface{}) error
	// readBE(dst interface{}) error
	ReadUBits(n uint8) (uint64, error)
	ReadSBits(n uint8) (int64, error)
	ReadFBits(n uint8) (float32, error)
	ReadByte() (byte, error)
	ReadBytes(n uint32) ([]byte, error)
	ReadAll() ([]byte, error)
	ReadBool() (bool, error)
	ReadUInt8() (uint8, error)
	ReadInt8() (int8, error)
	ReadUInt16() (uint16, error)
	ReadInt16() (int16, error)
	ReadUInt32() (uint32, error)
	ReadInt32() (int32, error)
	ReadUInt64() (uint64, error)
	ReadInt64() (int64, error)
	ReadFixed8() (swfprimitive.Fixed8, error)
	ReadFixed16() (swfprimitive.Fixed16, error)
	ReadFloat16() (swfprimitive.Float16, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
}

type swfreader struct {
	origin    io.Reader
	reader    *bitio.Reader
	bitsCount uint64
}

func NewReader(reader io.Reader) (Reader, error) {
	r := &swfreader{
		origin: reader,
		reader: bitio.NewReader(reader),
	}

	return r, nil
}

func (s *swfreader) PositionBytes() uint64 {
	remainder := s.bitsCount % byteBits

	return (s.bitsCount + remainder) / byteBits
}

func (s *swfreader) PositionBits() uint64 {
	return s.bitsCount
}

func (s *swfreader) UseZlibReader() error {
	rs, ok := s.origin.(io.ReadSeeker)

	if !ok {
		return fmt.Errorf("origin is not a ReadSeeker")
	}

	rs.Seek(SwfSignatureOffset, io.SeekStart)

	reader, err := zlib.NewReader(rs)

	if err != nil {
		return err
	}

	s.reader = bitio.NewReader(reader)

	return nil
}

func (s *swfreader) UseLzmaReader() error {
	rs, ok := s.origin.(io.ReadSeeker)

	if !ok {
		return fmt.Errorf("origin is not a ReadSeeker")
	}

	rs.Seek(SwfSignatureOffset, io.SeekStart)

	reader, err := lzma.NewReader(rs)

	if err != nil {
		return err
	}

	s.reader = bitio.NewReader(reader)

	return nil
}

func (s *swfreader) align() {
	skipped := s.reader.Align()

	s.bitsCount += uint64(skipped)
}

func (s *swfreader) readLE(dst interface{}) error {
	s.align()

	return binary.Read(s.reader, binary.LittleEndian, dst)
}

func (s *swfreader) readBE(dst interface{}) error {
	s.align()

	return binary.Read(s.reader, binary.BigEndian, dst)
}

func (s *swfreader) ReadUBits(n uint8) (uint64, error) {
	v, err := s.reader.ReadBits(n)

	if err != nil {
		return 0, err
	}

	s.bitsCount += uint64(n)

	return v, nil
}

func (s *swfreader) ReadSBits(n uint8) (int64, error) {
	bits, err := s.reader.ReadBits(n)

	shift := 32 - n

	if err != nil {
		return 0, err
	}

	s.bitsCount += uint64(n)

	return (int64(bits) << shift) >> shift, nil
}

func (s *swfreader) ReadFBits(n uint8) (float32, error) {
	v, err := s.ReadSBits(n)

	if err != nil {
		return 0, err
	}

	return float32(v) / 65536, nil
}

func (s *swfreader) ReadByte() (byte, error) {
	s.align()

	b, err := s.reader.ReadByte()

	if err != nil {
		return 0, err
	}

	s.bitsCount += byteBits

	return b, nil
}

func (s *swfreader) ReadBytes(n uint32) ([]byte, error) {
	s.align()

	bytes := make([]byte, n)

	count := 0

	for count < int(n) {
		c, err := s.reader.Read(bytes[count:n])

		if err != nil {
			return nil, err
		}

		s.bitsCount += uint64(c * byteBits)

		count += c
	}

	return bytes, nil
}

func (s *swfreader) ReadAll() ([]byte, error) {
	s.align()

	bytes, err := ioutil.ReadAll(s.reader)

	if err != nil {
		return nil, err
	}

	s.bitsCount += uint64(len(bytes) * byteBits)

	return bytes, err
}

func (s *swfreader) ReadBool() (bool, error) {
	var b bool

	err := s.readLE(&b)

	if err != nil {
		return false, err
	}

	s.bitsCount += 1

	return b, err
}

func (s *swfreader) ReadUInt8() (uint8, error) {
	var u uint8

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 8

	return u, err
}

// ReadInt8 reads an int8 from the reader.
func (s *swfreader) ReadInt8() (int8, error) {
	var i int8

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 8

	return i, err
}

// ReadUInt16 reads an uint16 from the reader.
func (s *swfreader) ReadUInt16() (uint16, error) {
	var u uint16

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 16

	return u, err
}

// ReadInt16 reads an int16 from the reader.
func (s *swfreader) ReadInt16() (int16, error) {
	var i int16

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 16

	return i, err
}

// ReadUInt32 reads an uint32 from the reader.
func (s *swfreader) ReadUInt32() (uint32, error) {
	var u uint32

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 32

	return u, err
}

// ReadInt32 reads an int32 from the reader.
func (s *swfreader) ReadInt32() (int32, error) {
	var i int32

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 32

	return i, err
}

// ReadUInt64 reads an uint64 from the reader.
func (s *swfreader) ReadUInt64() (uint64, error) {
	var u uint64

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 64

	return u, err
}

// ReadInt64 reads an int64 from the reader.
func (s *swfreader) ReadInt64() (int64, error) {
	var i int64

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 64

	return i, err
}

// ReadFixed8 reads a fixed16 from the reader.
func (s *swfreader) ReadFixed8() (swfprimitive.Fixed8, error) {
	var i swfprimitive.Fixed8

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 16

	return i, err
}

// ReadFixed16 reads a fixed16 from the reader.
func (s *swfreader) ReadFixed16() (swfprimitive.Fixed16, error) {
	var i swfprimitive.Fixed16

	err := s.readLE(&i)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 16

	return i, err
}

// ReadFloat16 reads a float16 from the reader.
func (s *swfreader) ReadFloat16() (swfprimitive.Float16, error) {
	var u swfprimitive.Float16

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 16

	return u, err
}

// ReadFloat32 reads a float32 from the reader.
func (s *swfreader) ReadFloat32() (float32, error) {
	var u float32

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 32

	return u, err
}

// ReadFloat64 reads a float64 from the reader.
func (s *swfreader) ReadFloat64() (float64, error) {
	var u float64

	err := s.readLE(&u)

	if err != nil {
		return 0, err
	}

	s.bitsCount += 64

	return u, err
}
