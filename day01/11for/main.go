package main

import "fmt"

//for循环

func main() {
	// 	 基本格式 for 初始语句;条件表达式;结束语句{
	//     循环体语句

	for i := 0; i < 10; i++ { //大部分用这种格式的for语句
		fmt.Println(i)
	}

	//for 变种写法
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种2
	// var i = 5
	// for ; i < 10; i++ {
	// 	println(i)
	// 	i++
	// }

	//死循环
	// for {
	// 	fmt.Println("123")
	// }

	//for range循环
	// s := "hello 李佳锦"
	// for i, v := range s { //i是索引
	// 	fmt.Printf("%d %c\n", i, v)
	// }

}
