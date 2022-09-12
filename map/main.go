package main

import "fmt"

func main() {

	type Stu struct { //定义一个结构体作为map中value的内容
		Name    string
		Age     int
		Address string
	}

	students := make(map[string]Stu, 10)
	stu1 := Stu{"kp", 22, "湖北省"}
	stu2 := Stu{"lxl", 21, "湖北省"}
	students["No1"] = stu1
	students["No2"] = stu2
	fmt.Println("students=", students)
	for n, info := range students {
		fmt.Println("学生的编号是：", n)
		fmt.Println("学生名字是：", info.Name)
		fmt.Println("学生年龄是：", info.Age)
		fmt.Println("学生地址是：", info.Address)
	}
}
