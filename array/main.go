package main

import "fmt"

func main() {
	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{2, 4, 5}
	fmt.Println("numArr02=", numArr02)

	//这里的[...]是固定写法
	var numArr03 = [...]int{1, 2, 3, 5, 6, 87, 6}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 100, 0: 200, 2: 300}
	fmt.Println("numArr04=", numArr04)
}
