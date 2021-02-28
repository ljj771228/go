package main

import "fmt"

//整型

func main() {
	//十进制数
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //十进制数转换为8进制数
	fmt.Printf("%x\n", i1) //十进制转换为16进制数
	fmt.Printf("%b\n", i1) //十进制转换为2进制数

	//定义8进制数077转换为十进制数是63
	i2 := 077
	fmt.Printf("%d\n", i2)
	//16进制数0xf转换为十进制数是15
	i3 := 0xf
	fmt.Printf("%d\n", i3)
	//%T大写的T，查看变量类型
	fmt.Printf("%T\n", i3)

	//声明int8类型变量
	i4 := int8(9) //明确指定int8类型，否则就是int类型
	fmt.Printf("%T\n", i4)

}
