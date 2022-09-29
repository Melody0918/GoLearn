package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string) {
	fmt.Printf("%v完成了减法运行,%v-%v=%v\n", name, this.Num1, this.Num2, this.Num1-this.Num2)
}

func TestReflect(b interface{}) {
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	rKd := rVal.Kind()
	if rKd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	numField := rVal.NumField()
	fmt.Printf("struct has %v Fields\n", numField)
	for i := 0; i < numField; i++ {
		fmt.Printf("Field%v为%v %v 值为%v\n", i, rType.Field(i).Name, rType.Field(i).Type, rVal.Field(i))
	}
	numMethod := rVal.NumMethod()
	fmt.Printf("struct has %v methods\n", numMethod)
	//使用反射机制调用结构体方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	rVal.Method(0).Call(params)
}

func main() {
	cal := Cal{
		Num1: 8,
		Num2: 3,
	}
	TestReflect(cal)
}
