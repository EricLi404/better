package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "124GZ0\u0000\u0000"
	b := "124GZ0"
	trimRight := strings.TrimRight(a, string([]byte{0}))
	fmt.Println(trimRight)
	fmt.Println(b)
}
