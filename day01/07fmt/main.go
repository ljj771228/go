package main

import "fmt"

//fmt 占位符
func main() {
	var n = 100
	//查看类型 %T
	fmt.Printf("%T\n", n)
	//查看变量值 %v
	fmt.Printf("%v\n", n)
	//查看二进制 %b
	fmt.Printf("%b\n", n)
	//查看十进制 %d
	fmt.Printf("%d\n", n)
	//查看八进制 %o
	fmt.Printf("%o\n", n)
	//查看十六进制 %x
	fmt.Printf("%x\n", n)
	var s = "李佳锦你好"
	//查看字符串 %s
	fmt.Printf("%s\n", s)

}
