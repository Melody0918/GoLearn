package main

import "fmt"

func main() {
	//channelpractice1
	//var x byte
	//var y byte = 'a'
	//fmt.Println("请输入转换为大写的char型(a,b,c,d,e中任选):")
	//fmt.Scanf("%c", &x)
	////fmt.Scanln(&x)
	//fmt.Printf("%c\n", y)
	//
	//switch x {
	//case 'a':
	//	fmt.Println("A")
	//case 'b':
	//	fmt.Println("B")
	//case 'c':
	//	fmt.Println("C")
	//case 'd':
	//	fmt.Println("D")
	//case 'e':
	//	fmt.Println("E")
	//default:
	//	fmt.Println("other")
	//}

	//practice2
	//var score float64
	//fmt.Println("请输入学生成绩：")
	//fmt.Scanln(&score)
	//
	//switch {
	//case score >= 60:
	//	fmt.Println("合格")
	//case score < 60.0:
	//	fmt.Println("不合格")
	//}

	//practice3
	//var month int
	//fmt.Println("请输入月份：")
	//fmt.Scanln(&month)

	//switch month {
	//case 3, 4, 5:
	//	fmt.Println("春季")
	//case 6, 7, 8:
	//	fmt.Println("夏季")
	//case 9, 10, 11:
	//	fmt.Println("秋季")
	//case 12, 1, 2:
	//	fmt.Println("冬季")
	//}

	//practice4
	var weekday string
	fmt.Println("请输入星期时间:")
	fmt.Scanln(&weekday)

	switch weekday {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("酸溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	}

}
