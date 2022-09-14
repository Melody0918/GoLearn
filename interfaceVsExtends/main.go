package main

import "fmt"

//Monkey结构体
type Monkey struct {
	Name string
}

//声明接口
type Bird interface {
	Fly()
}

func (monkey *Monkey) climb() {
	fmt.Println(monkey.Name + "生来会爬树")

}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//让LittleMonkey实现Bird
func (littlemonkey *LittleMonkey) Fly() {
	fmt.Println(littlemonkey.Name + "通过学习，会飞翔")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	var b Bird = &monkey //接口可以指向一个实现了该接口的自定义类型的变量（实例）
	b.Fly()
	monkey.climb()
	monkey.Fly()
}
