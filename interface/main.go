package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Camera struct {
	Name string
}
type Phone struct {
	Name string
}

//让Camera实现Usb接口的方法
func (c Camera) Start() {
	fmt.Printf("%s相机开始工作...\n", c.Name)
}
func (c Camera) Stop() {
	fmt.Printf("%s相机停止工作...\n", c.Name)
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Printf("%s手机开始工作...\n", p.Name)
}
func (p Phone) Stop() {
	fmt.Printf("%s手机停止工作...\n", p.Name)
}

type Computer struct {
}

//编写一个方法Working 方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working(usb Usb) {
	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	//测试
	//先创建结构体变量
	computer := Computer{}
	camera := Camera{"佳能"}
	phone := Phone{"iphone"}
	//关键点
	computer.Working(camera)
	computer.Working(phone)
}
