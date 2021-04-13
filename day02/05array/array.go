package main

import "fmt"

//数组

//存放元素的容器
//必须指定存放的元素的类型和容量
//数组的长度书数组类型的一部分

func main() {
	var a1 [3]bool //长度为3的数组[true fales true]
	var a2 [4]bool //[true true fales fales]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//数组的初始化
	//如果不初始化:默认元素都是零值（布尔值的默认flase，整型和浮点型都是0，字符串：“控制”）
	fmt.Println(a1, a2)
	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2.初始化方式2
	//根据[...]初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a10)
	//3.初始化方式3根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)
	//数组的遍历
	citys := [...]string{"北京", "上海", "南京"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//多维数组
	// [[1,2] [3,4] [5,6]]
	var a11 [3][2]int //一共有三个元素数组，每个元素都有长度为2的元素int型数组，数组里只能放相同类型
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
}
