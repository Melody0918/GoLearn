package main

import "fmt"

func main() {
	var slice1 []int = []int{100, 200, 300}
	fmt.Println("slice1=", slice1)
	fmt.Println("slice1容量为：", cap(slice1))
	fmt.Printf("slice1地址为：%p\n", &slice1)
	slice1 = append(slice1, 400, 500, 600)
	fmt.Println("slice1=", slice1)
	fmt.Printf("slice1地址为：%p\n", &slice1)
	fmt.Println("slice1容量为：", cap(slice1))

	var slice2 = make([]int, 1)
	fmt.Println(slice2)

	var slice3 []int
	fmt.Println(slice3)
}
