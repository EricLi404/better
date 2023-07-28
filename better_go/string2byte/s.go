package string2byte

import (
	"strings"
	"unsafe"
)

var ss = strings.Repeat("a", 1024)

func directToByteString(str string) []byte {
	return []byte(str)
}

// unsafeStrToBytes str转为byte
func unsafeStrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
