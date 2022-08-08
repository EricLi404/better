package s

import (
	"testing"
)

func Benchmark_sonicJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := sonicJson(testData)
		if err != nil {
			b.Fatal(err)
		}
		if string(bytes) != successString {
			b.Fatal()
		}
	}
}

func Benchmark_sonicJsonEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := sonicJsonEncoder(testData)
		if err != nil {
			b.Fatal(err)
		}
		if string(bytes) != successString {
			b.Fatal()
		}
	}
}

func Benchmark_nativeJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := nativeJson(testData)
		if err != nil {
			b.Fatal(err)
		}
		if string(bytes) != successString {
			b.Fatal()
		}
	}
}

func Benchmark_nativeJsonEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := nativeJsonEncoder(testData)
		if err != nil {
			b.Fatal(err)
		}
		if string(bytes) != successString {
			b.Fatal()
		}
	}
}
