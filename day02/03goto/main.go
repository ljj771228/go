package main

import "fmt"

//goto

func main() {
	//跳出多层for循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出外层的for循环
		}
	}
}
