



#### result

```text

go test -bench=. -cpu=1 -benchmem -benchtime=5s
goos: darwin
goarch: arm64
pkg: better_go/json/case_item

Benchmark__Unmarshal_ori                   61731             96828 ns/op           20464 B/op        746 allocs/op
Benchmark__Unmarshal_jsoniterator          88509             74886 ns/op           19808 B/op       1151 allocs/op
Benchmark__Unmarshal_easyjson             181471             33362 ns/op            3536 B/op        445 allocs/op
Benchmark__Unmarshal_sonic                 42972            138562 ns/op          117200 B/op       1346 allocs/op

Benchmark__Marshal_ori                     93595             63307 ns/op           24353 B/op        206 allocs/op
Benchmark__Marshal_jsoniterator           258008             23043 ns/op           25152 B/op        306 allocs/op
Benchmark__Marshal_easyjson               421423             14570 ns/op           13984 B/op        106 allocs/op
Benchmark__Marshal_sonic                   94976             67152 ns/op           29170 B/op        306 allocs/op
PASS
ok      better_go/json/case_item        54.678s


```



#### go test

```shell
go test -bench=. -cpu=1 -benchmem -benchtime=5s
```


#### easyjson
```shell

 easyjson -all common.go

```




