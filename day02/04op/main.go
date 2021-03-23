package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在等号的右边赋值 ==> a = a+1
	b-- //单独的语句，不能放在等号的右边赋值 ==> b = b-1
	fmt.Println(a)

	//关系运算符
	fmt.Println(a == b) //GO语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) //不等于
	fmt.Println(a >= b) //大于等于
	fmt.Println(a > b)  //大于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a < b)  //小于

	//逻辑运算符
	//如果年龄大于18岁，并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("老老实实上班吧")
	} else {
		fmt.Println("不用上班啦")
	}
	//如果年龄小于18岁，或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班啦")
	} else {
		fmt.Println("老老实实上班吧")
	}

	//not取反，原来为真就为假，原来为假，就为真
	isMarried := false
	fmt.Println(isMarried)  // false
	fmt.Println(!isMarried) //ture

	// 位运算符：针对的是二进制数
	// 5的二进制数标识：101
	// 2的二进制数标识： 10

	//&：按位与（两位均为1才为1）
	fmt.Println(5 & 2) //000
	//|: 按位或（两位有一个为1，就为1）
	fmt.Println(5 | 2) //111
	//^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //111
	//<<:将二进制位左移指定位数
	fmt.Println(5 << 1)  //1010  => 10
	fmt.Println(1 << 10) //100000000000  => 1024
	//<<:将二进制位右移指定的位数
	fmt.Println(5 >> 1) //10 ==> 2
	var m = int8(1)     //只能存8位
	fmt.Println(m << 7) //10000000000 超过了int8的8位则会舍掉超出部分

	//192.168.1.1
	//权限
	//赋值运算符

	var x int
	x = 10
	x += 1 //x = x + 1
	x -= 1 //x = x - 1
	x *= 2 //x = x * 2
	x /= 2 //x = x / 2
	x %= 2 //x = x % 2

	x <<= 2 // x = x << 2
	x &= 2  // x = x & 2
	x |= 2  // x = x | 2
	x ^= 2  // x = x ^ 2
	fmt.Println(x)

}
