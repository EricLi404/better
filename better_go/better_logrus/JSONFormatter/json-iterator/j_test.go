package json_iterator

import (
	"testing"
)

func Benchmark_jsoniterJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := jsoniterJson(testData)
		if err != nil {
			b.Fatal(err)
		}
		if string(bytes) != successString {
			b.Fatal()
		}
	}
}

func Benchmark_jsoniterJsonEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, err := jsoniterJsonEncoder(testData)
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
