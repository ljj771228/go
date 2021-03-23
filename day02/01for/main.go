package main

import "fmt"

//流程控制之跳出for循环

func main() {
	//当i=5时候就跳出for循环
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break //跳出for循环
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")

	//当i等于5时就跳出for循环，但不执行循环内部的打印语句
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //跳过5，继续下面的for语句 i++
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
