package string2byte

import (
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		directToByteString(ss)
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unsafeStrToBytes(ss)
	}
}
