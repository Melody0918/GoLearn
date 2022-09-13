package model

import "fmt"

type person struct {
	Name string //其他包不能直接访问
	age  int
	sal  float64
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确...")
		//有程序员给个默认值
		p.age = 18 //默认值
	}
}
func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal < 300000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不对...")
		p.sal = 0
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
