package main

import (
	"fmt"
)

type s struct {
	a string
	b string
	c int
}

func main() {
	// 初始化过程
	p := s{a: "我爱Go语言", b: "是", c: 2}
	fmt.Printf("变量p的地址%p", &p)
	fmt.Printf("slice的地址%v\n", p)


}
