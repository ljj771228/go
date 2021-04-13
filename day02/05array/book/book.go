// package main

// import "fmt"

// func main() {
// 	var a = [...]int{1, 3, 5, 7, 8}
// 	var sum int = 8

// 	for i := 0; i < len(a); i++ {
// 		sum += a[i]
// 		if i == len(a)-1 {
// 			fmt.Println("这个数组的和是:", sum)
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}

	sum := 0
	// for range(键值循环)：遍历数组、切片、字符串、map及通道（channel）
	// 通过for range遍历的返回值用_接收
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
}
