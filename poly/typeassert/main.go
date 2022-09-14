package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok
	//如何将a赋给一个Point变量？
	var b Point
	//b=a 不可以
	//b=a.(Point)//可以
	b = a.(Point) //类型断言
	fmt.Println(b)

	//类型断言的其他案例
	//var x interface{}
	//fmt.Printf("%T\n", x)
	//var f float32 = 1.1
	//x = f //空接口可以接收任意类型
	//fmt.Printf("%T\n", x)
	//// x=>float32[使用类型断言]
	//y := x.(float32) //类型断言
	//fmt.Printf("y的类型是%T 值是%v", y, y)

	//类型断言(带检测的)
	var x interface{}
	fmt.Printf("%T\n", x) //<nil>
	var f float32 = 1.1
	x = f                 //空接口可以接收任意类型
	fmt.Printf("%T\n", x) //float32
	// x=>float32[使用类型断言]
	y, ok := x.(float64) //类型断言
	//if y, ok := x.(float64); ok { //另外一种写法
	if ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T 值是%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
	fmt.Printf("y的类型是%T 值是%v", y, y) //y的类型是float32 值是1.1
}
