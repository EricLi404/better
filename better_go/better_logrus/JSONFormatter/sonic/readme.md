

## 结论

json-iterator encoder 效果并不好

换原生 iterator Marshal



```text

go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: linux
goarch: amd64
pkg: better_go/better_logrus/JSONFormatter/sonic
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz

Benchmark_sonicJson         	 8147275	     732.1 ns/op	     304 B/op	       5 allocs/op
Benchmark_sonicJsonEncoder  	 5413186	      1109 ns/op	     529 B/op	       8 allocs/op
Benchmark_nativeJson        	 5590626	      1073 ns/op	     240 B/op	       2 allocs/op
Benchmark_nativeJsonEncoder 	 5533387	      1088 ns/op	     144 B/op	       2 allocs/op
PASS
ok  	better_go/better_logrus/JSONFormatter/sonic	28.067s

```