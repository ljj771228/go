package main

import (
	"fmt"
)

//浮点数
func main() {
	//math.MaxFloat32 //float32位浮点数
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中小数都是float64位
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) //显式声明float32类型
}
