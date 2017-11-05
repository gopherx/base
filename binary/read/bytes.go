package read

import (
	"io"

	"github.com/gopherx/base/errors"
)

type BigEndian struct {
	r         io.Reader
	tmp       []byte
	Err       error
	ReadBytes int
}

func NewBigEndian(r io.Reader) *BigEndian {
	return &BigEndian{r, make([]byte, 12), nil, 0}
}

func (e *BigEndian) readTo(dest []byte) error {
	if e.Err != nil {
		return e.Err
	}

	rn, err := e.r.Read(dest)
	if err != nil {
		e.Err = err
	}

	if err == nil && rn != len(dest) {
		e.Err = errors.DataLoss(nil, "not enough data; read: ", rn, " wanted: ", len(dest))
	}

	e.ReadBytes += len(dest)
	return e.Err
}

func (e *BigEndian) read(n int) ([]byte, error) {
	err := e.readTo(e.tmp[0:n])
	return e.tmp, err
}

// Byte reads a byte from the buffer.
func (e *BigEndian) Byte() byte {
	b, err := e.read(1)
	if err != nil {
		return 0
	}

	return b[0]
}

// Uint16 reads an uin16 from the buffer.
func (e *BigEndian) Uint16() uint16 {
	b, err := e.read(2)
	if err != nil {
		return 0
	}

	b0 := uint16(b[0])
	b1 := uint16(b[1])
	return b0<<8 | b1
}

// Uint32 reads an uint32 from the buffer.
func (e *BigEndian) Uint32() uint32 {
	b, err := e.read(4)
	if err != nil {
		return 0
	}

	b0 := uint32(b[0])
	b1 := uint32(b[1])
	b2 := uint32(b[2])
	b3 := uint32(b[3])
	return b0<<24 | b1<<16 | b2<<8 | b3
}

// Uint64 reads an Uint64 from the buffer.
func (e *BigEndian) Uint64() uint64 {
	b, err := e.read(8)
	if err != nil {
		return 0
	}

	b0 := uint64(b[0])
	b1 := uint64(b[1])
	b2 := uint64(b[2])
	b3 := uint64(b[3])
	b4 := uint64(b[4])
	b5 := uint64(b[5])
	b6 := uint64(b[6])
	b7 := uint64(b[7])
	return b0<<56 | b1<<48 | b2<<40 | b3<<32 | b4<<24 | b5<<16 | b6<<8 | b7
}

// Int64 reads an Uint64 from the buffer.
func (e *BigEndian) Int64() int64 {
	b, err := e.read(8)
	if err != nil {
		return 0
	}

	b0 := int64(b[0])
	b1 := int64(b[1])
	b2 := int64(b[2])
	b3 := int64(b[3])
	b4 := int64(b[4])
	b5 := int64(b[5])
	b6 := int64(b[6])
	b7 := int64(b[7])
	return b0<<56 | b1<<48 | b2<<40 | b3<<32 | b4<<24 | b5<<16 | b6<<8 | b7
}

// Uint32x3 reads three uint32 from the buffer.
func (e *BigEndian) Uint32x3() (uint32, uint32, uint32) {
	b, err := e.read(12)
	if err != nil {
		return 0, 0, 0
	}

	b0 := uint32(b[0])
	b1 := uint32(b[1])
	b2 := uint32(b[2])
	b3 := uint32(b[3])

	b4 := uint32(b[4])
	b5 := uint32(b[5])
	b6 := uint32(b[6])
	b7 := uint32(b[7])

	b8 := uint32(b[8])
	b9 := uint32(b[9])
	b10 := uint32(b[10])
	b11 := uint32(b[11])

	return (b0<<24 | b1<<16 | b2<<8 | b3),
		(b4<<24 | b5<<16 | b6<<8 | b7),
		(b8<<24 | b9<<16 | b10<<8 | b11)
}

// Bytes reads n bytes from the reader.
func (e *BigEndian) Bytes(n int) []byte {
	tmp := make([]byte, n)
	e.readTo(tmp)
	return tmp
}

// Uint16 reads an uin16 from the buffer.
func Uint16(b []byte) uint16 {
	b0 := uint16(b[0])
	b1 := uint16(b[1])
	return b0<<8 | b1
}

// Uint32 reads an uint32 from the buffer.
func Uint32(b []byte) uint32 {
	b0 := uint32(b[0])
	b1 := uint32(b[1])
	b2 := uint32(b[2])
	b3 := uint32(b[3])
	return b0<<24 | b1<<16 | b2<<8 | b3
}

// Uint64 reads an Uint64 from the buffer.
func Uint64(b []byte) uint64 {
	b0 := uint64(b[0])
	b1 := uint64(b[1])
	b2 := uint64(b[2])
	b3 := uint64(b[3])
	b4 := uint64(b[4])
	b5 := uint64(b[5])
	b6 := uint64(b[6])
	b7 := uint64(b[7])
	return b0<<56 | b1<<48 | b2<<40 | b3<<32 | b4<<24 | b5<<16 | b6<<8 | b7
}

// Int64 reads an Uint64 from the buffer.
func Int64(b []byte) int64 {
	b0 := int64(b[0])
	b1 := int64(b[1])
	b2 := int64(b[2])
	b3 := int64(b[3])
	b4 := int64(b[4])
	b5 := int64(b[5])
	b6 := int64(b[6])
	b7 := int64(b[7])
	return b0<<56 | b1<<48 | b2<<40 | b3<<32 | b4<<24 | b5<<16 | b6<<8 | b7
}

// Uint32x3 reads three uint32 from the buffer.
func Uint32x3(b []byte) (uint32, uint32, uint32) {
	b0 := uint32(b[0])
	b1 := uint32(b[1])
	b2 := uint32(b[2])
	b3 := uint32(b[3])

	b4 := uint32(b[4])
	b5 := uint32(b[5])
	b6 := uint32(b[6])
	b7 := uint32(b[7])

	b8 := uint32(b[8])
	b9 := uint32(b[9])
	b10 := uint32(b[10])
	b11 := uint32(b[11])

	return (b0<<24 | b1<<16 | b2<<8 | b3),
		(b4<<24 | b5<<16 | b6<<8 | b7),
		(b8<<24 | b9<<16 | b10<<8 | b11)
}
