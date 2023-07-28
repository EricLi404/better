package main

import "fmt"
import _ "unsafe"

// 这样就可以把当前的 函数 z，映射为 math 库中的 zero 函数，这样就可以用那些未导出的函数了
// 注意两个函数签名得一样
// 这个操作同样适用于变量
//
//go:linkname z math.zero
func z(x uint64) uint64

func main() {
	fmt.Println(z(5))
}
