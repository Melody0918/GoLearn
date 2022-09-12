package main //一个Go文件需要归属于一个包

import "fmt" // 引入一个包

func main() { //func是一个关键字，表示一个函数；main是函数名，是一个主函数，即我们程序的入口
	fmt.Println("hello,world!")
	fmt.Println(getSumAndSub(3, 4)) // 返回函数
	fmt.Println("张三\n赵六李四\r王五")
}

// Go语言指针的使用特点
func testPtr(num *int) {
	*num = 20
}

// 写一个函数，实现同时返回 和，差
// go函数支持返回多个值
func getSumAndSub(n1 int, n2 int) (int, int) {

	sum := n1 + n2 //go语言后面不要带分号
	sub := n1 - n2
	return sum, sub
}
