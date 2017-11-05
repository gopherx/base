package format

import (
	"fmt"

	"strconv"
)

func HexU8(v byte) string {
	return strconv.FormatUint(uint64(v), 16)
}

func SHex(b []byte) string {
	s := "["
	sep := ""
	for _, v := range b {
		s += sep
		s += HexU8(v)
		sep = " "
	}
	s += "]"
	return s
}

func Hex4C(b []byte) string {
	s := ""
	sep := "  1: "

	for i, v := range b {
		s += sep
		vs := HexU8(v)
		if len(vs) == 1 {
			s += "0"
		}
		s += vs

		if (i+1)%4 == 0 {
			s += "\n"
			sep = fmt.Sprintf("%3d: ", i)
		} else {
			sep = " "
		}

	}
	return s
}

func OutHex4C(b []byte) {
	fmt.Print(Hex4C(b))
}
