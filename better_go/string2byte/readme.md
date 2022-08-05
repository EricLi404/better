

# 优化逻辑

string 与 []byte 在结构上仅差一个 Cap，所以可以直接 unsafe 转.


```go
package string2byte

type StringHeader struct {
	Data uintptr
	Len  int
}
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

```


```text

go test -bench=. -cpu=1 -benchmem

goos: darwin
goarch: arm64
pkg: better_go/string2byte


BenchmarkDirectToString         514813948                2.309 ns/op           0 B/op          0 allocs/op
BenchmarkUnsafeToString         1000000000               0.3145 ns/op          0 B/op          0 allocs/op


BenchmarkDirectToBytes          515457274                2.342 ns/op           0 B/op          0 allocs/op
BenchmarkUnsafeToBytes1         1000000000               0.3184 ns/op          0 B/op          0 allocs/op
BenchmarkUnsafeToBytes2         1000000000               0.3145 ns/op          0 B/op          0 allocs/op


PASS
ok      better_go/string2byte   4.368s

```