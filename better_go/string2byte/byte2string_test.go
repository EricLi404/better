package string2byte

import (
	"fmt"
	"testing"
	"unsafe"
)

func BenchmarkDirectToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		directToString(by)
	}
}

func BenchmarkUnsafeToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString(by)
	}
}

func Test112(t *testing.T) {
	fmt.Println(string(by))
	xx := toString(by)
	fmt.Println(xx)
	by[1] = byte('m')
	fmt.Println(xx)
}

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func Test666(t *testing.T) {
	x := BytesToString(by)
	fmt.Println(x)
	by[1] = byte('m')
	fmt.Println(x)
}
