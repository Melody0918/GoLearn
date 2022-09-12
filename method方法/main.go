package main

import "fmt"

type Person struct {
	Name string
}

//给Person类型绑定一个方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}
func main() {
	var p Person
	p.Name = "james"
	p.test()
}
