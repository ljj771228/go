package main

import "fmt"

//字符串
func main() {
	// \ 本来是具有特殊含义的，我应该告诉程序我写的 \ 就是一个单纯的 \
	path := "\"D:\\ESXI+IKUAI+LEDE\""
	fmt.Println(path)

	s := "i'm ok !"
	fmt.Println(s)
	//多行字符串
	s2 := `
	诗情薄
	啦啦啦
	嘎嘎嘎
	哒哒哒	
	`
	fmt.Println(s2)

	s3 := `"D:\ESXI+IKUAI+LEDE\"`
	fmt.Println(s3)
	//字符串长度
	fmt.Println(len(s3))
	//字符串拼接
	name := "八嘎"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
}
