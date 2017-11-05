package write

import (
	"math/rand"
	"reflect"
	"testing"
	//"github.com/gopherx/base/binary/format"
)

func bytes(b ...byte) []byte {
	return b
}

func TestBigEndian(t *testing.T) {
	d := make([]byte, 28)
	w := BigEndian{d, 0, nil}
	w.Uint16(0xF00D)
	w.Uint32(0xBAADF00D)
	w.Uint64(0xAAAABBBBCCCCDDDD)
	w.Bytes([]byte{
		0xAB, 0xBC, 0xCD, 0xDE,
		0xAB, 0xBC, 0xCD, 0xDE,
		0xAB, 0xBC, 0xCD, 0xDE},
	)

	if w.Err != nil {
		t.Fatal(w)
	}

	want := []byte{
		0xF0, 0x0D,
		0xBA, 0xAD, 0xF0, 0x0D,
		0xAA, 0xAA, 0xBB, 0xBB,
		0xCC, 0xCC, 0xDD, 0xDD,
		0xAB, 0xBC, 0xCD, 0xDE,
		0xAB, 0xBC, 0xCD, 0xDE,
		0xAB, 0xBC, 0xCD, 0xDE,
		0x00, 0x00,
	}
	if !reflect.DeepEqual(w.Dest, want) {
		t.Fatal(w.Dest, want)
	}

	w.Uint32(0x11223344)
	if w.Err == nil {
		t.Fatal("last write should fail the writer")
	}

	//...try again with all ops to ensure we don't panic
	w.Uint16(0x00)
	w.Uint32(0x00)
	w.Uint64(0x00)
	w.Bytes([]byte{0x00})
}

func SliceAndWrite(b []byte, offset int, l int) {
	s := b[offset : offset+l]

	for i := 0; i < len(s); i++ {
		s[i] = byte(i)
	}
}

func BenchmarkSliceAndWrite(b *testing.B) {
	buff := make([]byte, 4096)
	rnd := rand.New(rand.NewSource(123234345))

	offset := 0
	for i := 0; i < b.N; i++ {
		l := 1 + int(rnd.Int31n(8))
		if offset+l > len(buff) {
			offset = 0
		}

		SliceAndWrite(buff, offset, l)

		offset += l
	}
}

func OffsetWrite(b []byte, offset int, l int) {
	if offset+l > len(b) {
		panic("ooooops")
	}

	for i := 0; i < l; i++ {
		b[offset+i] = byte(i)
	}
}

func BenchmarkOffsetAndWrite(b *testing.B) {
	buff := make([]byte, 4096)
	rnd := rand.New(rand.NewSource(123234345))

	offset := 0
	for i := 0; i < b.N; i++ {
		l := int(rnd.Int31n(8))
		if offset+l > len(buff) {
			offset = 0
		}

		OffsetWrite(buff, offset, l)

		offset += l
	}
}
