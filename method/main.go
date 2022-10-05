package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
}
type Circle struct {
	radius float64
}

//给Person类型绑定一个方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

// 为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area() float64 {
	//因为c是指针，因此我们标准的访问其字段的方式是((*c).radius)
	//return math.Pi * (*c).radius * (*c).radius
	return math.Pi * c.radius * c.radius //go 的设计者为了程序员使用方便，底层会对 c.radius进行处理，会加上*号
}

func main() {
	var p Person
	var c Circle
	p.Name = "james"
	c.radius = 4.0
	p.test()
	fmt.Println((&c).area())
	//编译器底层做了优化 (&c).area() 等价 c.area()

}
