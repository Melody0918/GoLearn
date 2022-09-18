package model

import "fmt"

//声明一个Customer结构体，表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}

//返回用户的信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t%v", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}

//修改用户信息
func (this *Customer) Update() {
	//fmt.Printf("客户原Id%v改为：", this.Id)
	//fmt.Scanln(&this.Id)
	fmt.Printf("客户原姓名(%v)改为：", this.Name)
	fmt.Scanln(&this.Name)
	fmt.Printf("客户原性别(%v)改为：", this.Gender)
	fmt.Scanln(&this.Gender)
	fmt.Printf("客户原年龄(%v)改为：", this.Age)
	fmt.Scanln(&this.Age)
	fmt.Printf("客户原电话(%v)改为：", this.Phone)
	fmt.Scanln(&this.Phone)
	fmt.Printf("客户原邮箱(%v)改为：", this.Email)
	fmt.Scanln(&this.Email)
}
