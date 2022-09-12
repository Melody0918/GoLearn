package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 = 3.4
	var c1 = '\n'
	var b1 = false
	var d1 float64 //输出默认值
	var e1 int
	var isMarry bool
	var str string
	fmt.Printf("b的数据类型为：%T\n", b1)
	fmt.Println("b占用的空间为：", unsafe.Sizeof(b1))
	fmt.Println("c1=", c1)
	fmt.Printf("d1=%v,e1=%d,isMarry=%v,str=%v\n", d1, e1, isMarry, str) // %v表示按照变量的值输出
	fmt.Printf("num1的数据类型是%T\n", num1)
	fmt.Println("姓名\n性别\n籍贯\n住址")
	loveGo()
}

func loveGo() {
	fmt.Println("\t*\t\t\t\t*\n")
	fmt.Println("*\t\t*   I Love Golang  *\t\t*\n")
	fmt.Println("\t*\t\t\t\t*\n")
	fmt.Println("\t      *\t\t\t*\n")
	fmt.Println("\t\t\t*")
}
