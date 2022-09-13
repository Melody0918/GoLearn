package main

import (
	"GoProject/factoryModel/model"
	"fmt"
)

func main() {
	//var student model.Student
	//student.Name = "kp"
	//student.Age = 22

	//当student结构体是首字母小写，我们可以通过工厂模式来解决
	var student = model.Newstudent("kp~", 22)

	fmt.Println("stu=", *student) //stu= {kp~ 22}
	fmt.Println("name=", student.Name, "age=", student.GetScore())
}
