

## 结论

json-iterator encoder 效果并不好

换原生 iterator Marshal



```text

go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: darwin
goarch: arm64
pkg: better_go/better_logrus/JSONFormatter/json-iterator
Benchmark_jsoniterJson          22721110               262.2 ns/op           248 B/op          3 allocs/op
Benchmark_jsoniterJsonEncoder   17202775               350.6 ns/op           768 B/op          6 allocs/op
Benchmark_nativeJson            20376416               295.6 ns/op           240 B/op          2 allocs/op
Benchmark_nativeJsonEncoder     19283504               306.3 ns/op           144 B/op          2 allocs/op
PASS
ok      better_go/better_logrus/JSONFormatter/json-iterator     25.258s


go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: linux
goarch: amd64
pkg: better_go/better_logrus/JSONFormatter/json-iterator
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
Benchmark_jsoniterJson        	 5778958	      1031 ns/op	     248 B/op	       3 allocs/op
Benchmark_jsoniterJsonEncoder 	 4598872	      1318 ns/op	     768 B/op	       6 allocs/op
Benchmark_nativeJson          	 5815002	      1041 ns/op	     240 B/op	       2 allocs/op
Benchmark_nativeJsonEncoder   	 5569124	      1069 ns/op	     144 B/op	       2 allocs/op
PASS
ok  	better_go/better_logrus/JSONFormatter/json-iterator	28.538s

```