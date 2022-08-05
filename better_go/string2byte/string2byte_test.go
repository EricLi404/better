package string2byte

import "testing"

func BenchmarkDirectToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		directToBytes(s)
	}
}

func BenchmarkUnsafeToBytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes1(s)
	}
}

func BenchmarkUnsafeToBytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeToBytes2(s)
	}
}
