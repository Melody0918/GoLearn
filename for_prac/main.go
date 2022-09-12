package main

import "fmt"

func main() {

	//practice1
	//var i int = 1
	//var count int
	//var sum int
	//for i <= 100 {
	//	if i%9 == 0 {
	//		fmt.Println(i)
	//		count++
	//		sum += i
	//	}
	//	i++
	//}
	//fmt.Printf("1~100之间所有是9的倍数的整数个数为:%d 总和为:%d", count, sum)
	////1~100之间所有是9的倍数的整数个数为:11 总和为:594

	//practice2
	var x int = 0
	for x <= 6 {
		fmt.Printf("%d + %d = 6\n", x, 6-x)
		x++
	}
}
