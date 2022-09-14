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

func (p Phone) call() {
	fmt.Println(p.Name + "手机打电话..")
}

//定义Computer结构体
type Computer struct {
	Name string
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向phone结构体变量，还需要调用call方法
	//类型断言+检测机制..[注意体会！！！]
	if phone, ok := usb.(Phone); ok {
		phone.call()
	}
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera结构体变量
	//这里就体现出多态数组,利用了接口的多态特点，使得数组中可以存放多个不同的数据类型
	var usbArr [3]Usb
	usbArr[0] = Phone{"iphone"}
	usbArr[1] = Phone{"HuaWei"}
	usbArr[2] = Camera{"佳能"}

	//遍历usbArr
	//phone还有一个特有的方法call()，请遍历Usb数组，如果是phone变量，
	//除了调用Usb接口声明的方法外，还需要调用phone特有方法call.=>类型断言
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}

	fmt.Println(usbArr) //不放东西时为:[<nil> <nil> <nil>]
}
