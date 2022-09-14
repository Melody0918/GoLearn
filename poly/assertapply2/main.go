package main

import "fmt"

//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {

	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index+1, x)
		case float32, float64:
			fmt.Printf("第%v个参数是浮点数类型，值是%v\n", index+1, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, x)
		case byte:
			fmt.Printf("第%v个参数是byte类型，值是%v\n", index+1, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index+1, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", index+1, x)
		default:
			fmt.Printf("第%v个参数是 类型不确定，值是%v\n", index+1, x)
		}
	}
}

type Student struct {
}

func main() {

	var n1 float32 = 1.2
	var n2 float64 = 5.78
	var n3 int32 = 30
	var name string = "kp"
	var stu1 Student
	var stu2 *Student = &Student{}
	stu3 := &Student{}
	fmt.Printf("*Student类型为%T\n", stu2)
	address := "南昌"
	n4 := 32
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2, stu3)
}
