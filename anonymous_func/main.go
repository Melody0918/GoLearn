package main

import "fmt"

func main() {
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能用一次

	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res=", res)
}
