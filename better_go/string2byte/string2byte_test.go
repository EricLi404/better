package string2byte

import (
	"fmt"
	"testing"
	"unsafe"
)

func toBytes22(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

type A struct {
	a int
	b int
}

func Test3434(t *testing.T) {
	b := [2]int{2, 3}
	c := *(*A)(unsafe.Pointer(&b))
	fmt.Println(c)

}

func Test11(t *testing.T) {
	xx := unsafeToBytes2(s)
	xx[1] = byte('d')
	fmt.Println(string(xx))
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func Test2123(t *testing.T) {
	xx := StringToBytes(s)
	fmt.Println(string(xx))
	xx[1] = byte('m')
	fmt.Println(string(xx))
}

func BenchmarkDirectToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toBytes22(s)
	}
}

func BenchmarkUnsafeToBytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes1(s)
	}
}

func BenchmarkUnsafeToBytes11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes11(s)
	}
}

func BenchmarkUnsafeToBytes12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes12(s)
	}
}

func BenchmarkUnsafeToBytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes2(s)
	}
}
