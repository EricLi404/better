



```text
1 -- 直接转
2 -- reflect
3 -- unsafe



 go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: darwin
goarch: arm64
pkg: better_go/string2byte/1.20
Benchmark__StringToByte1        64971511                93.17 ns/op         1024 B/op          1 allocs/op
Benchmark__StringToByte2        1000000000               0.3303 ns/op          0 B/op          0 allocs/op
Benchmark__StringToByte3        1000000000               0.4946 ns/op          0 B/op          0 allocs/op
Benchmark__ByteToString1        69014137                91.94 ns/op         1024 B/op          1 allocs/op
Benchmark__ByteToString2        1000000000               0.3309 ns/op          0 B/op          0 allocs/op
Benchmark__ByteToString3        1000000000               0.4996 ns/op          0 B/op          0 allocs/op
PASS
ok      better_go/string2byte/1.20      14.800s

```