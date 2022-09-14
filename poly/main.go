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

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera结构体变量
	//这里就体现出多态数组,利用了接口的多态特点，使得数组中可以存放多个不同的数据类型
	var usbArr [3]Usb
	usbArr[0] = Phone{
		"iphone",
	}
	usbArr[1] = Phone{
		"HuaWei",
	}
	usbArr[2] = Camera{
		"佳能",
	}
	fmt.Println(usbArr) //不放东西时为:[<nil> <nil> <nil>]
}
