package main

import "fmt"

//if条件判断
func main() {
	age := 19
	// if age > 18 { //如果age > 18就执行这个{}中的代码，
	// 	fmt.Println("开业了开业了")
	// } else { // 否则就执行这个{}中的代码
	// 	fmt.Println("去写作业了")
	// }

	//多个条件判断
	if age > 35 {
		fmt.Println("你是中年了")
	} else if age > 18 {
		fmt.Println("你是小年轻")
	} else {
		fmt.Println("好好学习")
	}
}
