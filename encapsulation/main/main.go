package main

import (
	"GoProject/encapsulation/model"
	"fmt"
)

func main() {
	person := model.NewPerson("kp")
	person.SetAge(22)
	person.SetSal(30000)
	fmt.Println(*person) //{kp 22 30000}
	fmt.Println("姓名：", person.Name, "年龄：", person.GetAge(), "薪水：", person.GetSal())
} //姓名： kp 年龄： 22 薪水： 30000
