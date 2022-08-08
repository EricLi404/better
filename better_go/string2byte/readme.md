

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

BenchmarkDirectToString         485622948                2.266 ns/op           0 B/op          0 allocs/op
BenchmarkUnsafeToString         1000000000               0.3120 ns/op          0 B/op          0 allocs/op

BenchmarkDirectToBytes          512121726                2.325 ns/op           0 B/op          0 allocs/op
BenchmarkUnsafeToBytes1         1000000000               0.3128 ns/op          0 B/op          0 allocs/op
BenchmarkUnsafeToBytes11        1000000000               0.3141 ns/op          0 B/op          0 allocs/op
BenchmarkUnsafeToBytes12        1000000000               0.3137 ns/op          0 B/op          0 allocs/op
BenchmarkUnsafeToBytes2         1000000000               0.3154 ns/op          0 B/op          0 allocs/op
PASS
ok      better_go/string2byte   5.058s



go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: linux
goarch: amd64
pkg: better_go/string2byte
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz

BenchmarkDirectToString  	885053401	          6.783 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeToString  	1000000000	         0.4197 ns/op	       0 B/op	       0 allocs/op

BenchmarkDirectToBytes   	594386966	          10.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeToBytes1  	1000000000	         0.4190 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeToBytes11 	1000000000	         0.4392 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeToBytes12 	1000000000	         0.4210 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnsafeToBytes2  	1000000000	         0.4203 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	better_go/string2byte	16.041s
```