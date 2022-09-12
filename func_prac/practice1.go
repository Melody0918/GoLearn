package main

import "fmt"

func main() {
	var year int
	var month int
	for {
		fmt.Print("请输入年：")
		fmt.Scanln(&year)
		fmt.Print("请输入月：")
		fmt.Scanln(&month)
		if year%4 == 0 && year%100 != 0 {
			switch month {
			case 2:
				fmt.Println("您输入的日期是：", 29)
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Println("您输入的日期是：", 31)
			case 4, 6, 9, 11:
				fmt.Println("您输入的日期是：", 30)
			default:
				fmt.Println("日期输入错误！")
			}
		} else if year > 0 {
			switch month {
			case 2:
				fmt.Println("您输入的日期是：", 28)
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Println("您输入的日期是：", 31)
			case 4, 6, 9, 11:
				fmt.Println("您输入的日期是：", 30)
			default:
				fmt.Println("日期输入错误！")
			}
		} else {
			fmt.Println("年份输入错误！")
		}
	}
}
