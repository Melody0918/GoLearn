package main

import "fmt"

func main() {
	//从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入0时结束程序
	//var positiveCount int = 0
	//var negativeCount int = 0
	//var num int
	//for {
	//	fmt.Println("请输入一个整数：")
	//	fmt.Scanln(&num)
	//
	//	if num == 0 {
	//		break
	//	}
	//
	//	if num > 0 {
	//		positiveCount++
	//		continue
	//	} else {
	//		negativeCount++
	//		continue
	//	}
	//}
	//fmt.Printf("正数的个数是：%v,负数的个数是：%v", positiveCount, negativeCount)

	//label2:
	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue label1
			}
			fmt.Println("j=", j)
		}
	}
}
