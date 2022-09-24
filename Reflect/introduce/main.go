package main

import (
	"fmt"
	"reflect"
)

func reflectTest1(b interface{}) {
	//通过反射获取传入的变量的type、kind、值
	//1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//2.获取到reflect.Value
	rval := reflect.ValueOf(b)
	//n1 := 10 + rval //int类型无法与reflect.Value类型相加
	n1 := 10 + rval.Int()
	fmt.Println(n1)
	fmt.Printf("rval=%v   rval type=%T\n", rval, rval)

	//下面我们将 rval 转换成 interface{}
	iv := rval.Interface()
	//将interface{} 通过类型断言转成需要的类型
	num2 := iv.(int)
	fmt.Printf("num2=%v   num2 type=%T\n", num2, num2)
}

func reflectTest2(b interface{}) {
	//通过反射获取传入的变量的type、kind、值
	//1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v rType Kind=%v rType Name=%v\n", rType, rType.Kind(), rType.Name())
	//2.获取到reflect.Value
	rval := reflect.ValueOf(b)
	fmt.Printf("rval=%v   rval type=%T\n	  rval kind=%v\n", rval, rval, rval.Kind())

	//下面我们将 rval 转换成 interface{}
	iv := rval.Interface()
	fmt.Printf("iv=%v   iv type=%T\n", iv, iv) //iv={tom 22} iv type=main.Student
	//将interface{} 通过类型断言转成需要的类型
	//可以参考类型断言最佳实践2，用switch的断言形式做的更灵活
	//这里就简单使用了带检测的类型断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\nstu.Age=%v\n", stu.Name, stu.Age)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1.先定义一个int
	var num int = 100
	reflectTest1(num)
	fmt.Println("---------------------------------")
	//2.定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  22,
	}
	//演示对(结构体类型、interface{}、reflect.Value)进行反射的基本操作
	reflectTest2(stu)
}
