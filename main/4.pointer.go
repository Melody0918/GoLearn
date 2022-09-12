package main

import "fmt"

func main() {

	//演示golang中的指针数据类型

	//基本数据类型在内存中的布局
	var i int = 10
	var x float64 = 23.5341
	// i的地址是什么，&i地址符
	fmt.Println("i的地址=", &i)

	//下面的var ptr *int = &i
	//1.ptr 是一个指针变量
	//2.ptr 的类型*int
	//3.ptr 本身的值是&i
	var ptr *int = &i
	var ptr1 *float64 = &x
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr1 = %v\n", ptr1)
	fmt.Println("ptr的地址是:", &ptr)
	fmt.Printf("ptr 指向的值是:%v\n", *ptr)
	*ptr = 66
	fmt.Printf("i的值变为:%v,i的数据类型为:%T", i, *ptr)
}
