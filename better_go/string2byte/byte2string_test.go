package string2byte

import "testing"

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
