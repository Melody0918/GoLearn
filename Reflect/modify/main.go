package main

import (
	"fmt"
	"reflect"
)

//通过反射修改变量值
func reflect1(b interface{}) {
	//2.获取到reflect.Value
	rval := reflect.ValueOf(b)
	//查看rval的Kind是什么
	fmt.Printf("rval kind=%v\nrval type=%v\n", rval.Kind(), rval.Type()) //ptr指针类型
	rval.Elem().SetInt(30)
}

func main() {
	var num int = 20
	reflect1(&num) //传入地址
	fmt.Println("num=", num)

	//reflect.Value.Elem()用于获取指针指向变量，类似
	var num1 = 10
	var b *int = &num1
	*b = 3
	fmt.Println(num1)
}
