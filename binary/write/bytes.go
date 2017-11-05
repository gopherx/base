package write

import (
	"github.com/gopherx/base/errors"
)

type BigEndian struct {
	Dest   []byte
	Offset int
	Err    error
}

func (b *BigEndian) fail(op string, v interface{}) {
	b.Err = errors.OutOfRange(nil, op, v)
}

func (b *BigEndian) Byte(v byte) {
	if b.Err != nil {
		return
	}

	if b.Offset+1 >= len(b.Dest) {
		b.fail("Byte:", v)
		return
	}

	b.Dest[b.Offset] = v
	b.Offset += 1
}

func (b *BigEndian) Uint16(v uint16) {
	if b.Err != nil {
		return
	}

	if b.Offset+2 >= len(b.Dest) {
		b.fail("Uint16:", v)
		return
	}

	b0 := byte(v >> 8)
	b1 := byte(v)

	dest := b.Dest
	offset := b.Offset
	dest[offset] = b0
	dest[offset+1] = b1

	b.Offset += 2
}

func (b *BigEndian) Uint16At(offset int, v uint16) {
	if b.Err != nil {
		return
	}

	if offset+2 >= len(b.Dest) {
		b.fail("Uint16At:", v)
		return
	}

	b0 := byte(v >> 8)
	b1 := byte(v)

	dest := b.Dest
	dest[offset] = b0
	dest[offset+1] = b1
}

func (b *BigEndian) Uint32(v uint32) {
	if b.Err != nil {
		return
	}

	if b.Offset+4 >= len(b.Dest) {
		b.fail("Uint32:", v)
		return
	}

	b0 := byte(v >> 24)
	b1 := byte(v >> 16)
	b2 := byte(v >> 8)
	b3 := byte(v)

	dest := b.Dest
	offset := b.Offset
	dest[offset] = b0
	dest[offset+1] = b1
	dest[offset+2] = b2
	dest[offset+3] = b3

	b.Offset += 4
}

func (b *BigEndian) Uint64(v uint64) {
	if b.Err != nil {
		return
	}

	if b.Offset+8 >= len(b.Dest) {
		b.fail("Uint64:", v)
		return
	}

	b0 := byte(v >> 56)
	b1 := byte(v >> 48)
	b2 := byte(v >> 40)
	b3 := byte(v >> 32)
	b4 := byte(v >> 24)
	b5 := byte(v >> 16)
	b6 := byte(v >> 8)
	b7 := byte(v)

	dest := b.Dest
	offset := b.Offset
	dest[offset] = b0
	dest[offset+1] = b1
	dest[offset+2] = b2
	dest[offset+3] = b3
	dest[offset+4] = b4
	dest[offset+5] = b5
	dest[offset+6] = b6
	dest[offset+7] = b7

	b.Offset += 8
}

func (b *BigEndian) Bytes(bytes []byte) {
	if b.Err != nil {
		return
	}

	if b.Offset+len(bytes) >= len(b.Dest) {
		b.fail("Bytes:", len(bytes))
		return
	}

	offset := b.Offset
	for i, v := range bytes {
		b.Dest[offset+i] = v
	}

	b.Offset += len(bytes)
}
