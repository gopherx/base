package read

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
