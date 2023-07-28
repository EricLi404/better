package __20

import (
	"bytes"
	"strings"
	"testing"
	"unsafe"
)

var str = strings.Repeat("x", 1024)
var by = bytes.Repeat([]byte{'x'}, 1024)

// 直接转
func Benchmark__StringToByte1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}

// reflect
func Benchmark__StringToByte2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := (*[2]uintptr)(unsafe.Pointer(&str))
		bb := [3]uintptr{x[0], x[1], x[1]}
		_ = *(*[]byte)(unsafe.Pointer(&bb))
	}
}

// go1.20 unsafe
func Benchmark__StringToByte3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.Slice(unsafe.StringData(str), len(str))
	}
}

// 直接转
func Benchmark__ByteToString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(by)
	}
}

// reflect
func Benchmark__ByteToString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&by))
	}
}

// go1.20 unsafe
func Benchmark__ByteToString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unsafe.String(unsafe.SliceData(by), len(by))
	}
}
