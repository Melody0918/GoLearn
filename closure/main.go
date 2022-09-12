package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

//闭包的最佳实现
func makeSuffix(suffix string) func(string) string {
	var str string = suffix
	return func(filename string) string {
		if !strings.HasSuffix(filename, str) {
			filename = filename + str
		}
		return filename
	}
}

func main() {

	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
	ms := makeSuffix(".jpg")
	fmt.Println(ms("kp"))
	fmt.Println(ms("123.jpg"))
}
