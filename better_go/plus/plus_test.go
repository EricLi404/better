package plus

import "testing"

func Benchmark_1(b *testing.B) {
	ii := 0
	for i := 0; i < b.N; i++ {
		ii += 1
	}
}

func Benchmark_2(b *testing.B) {
	ii := 0
	for i := 0; i < b.N; i++ {
		ii = ii + 1
	}
}
