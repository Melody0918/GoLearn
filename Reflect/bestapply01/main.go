package main

import (
	"fmt"
	"reflect"
)

//定义了一个Monster结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (this Monster) Print() {
	fmt.Println("---------start----------")
	fmt.Println(this)
	fmt.Println("---------end------------")
}

//返回两个数的和
func (this Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//
func (this Monster) Set(name string, age int, score float32, sex string) {
	this.Name = name
	this.Age = age
	this.Score = score
	this.Sex = sex
}
func TestStruct(a interface{}) {
	aType := reflect.TypeOf(a)
	aVal := reflect.ValueOf(a)
	aKd := aVal.Kind()
	if aKd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有多少个字段
	num := aVal.NumField()
	fmt.Printf("struct has %v fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %v:值为%v\n", i, aVal.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := aType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %v:tag为%v\n", i, tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := aVal.NumMethod()
	fmt.Printf("struct has %v methods\n", numOfMethod)

	//方法的排序默认是按照函数名排序(根据ASCII码)
	aVal.Method(1).Call(nil) //获取到第二个方法，调用它
	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := aVal.Method(0).Call(params) //传入的参数是 []reflect.Value
	fmt.Println("res=", res[0].Int())  //返回结果，结果是 []reflect.Value,因此转化为int

}

func main() {
	//创建一个Monster实例
	var monster Monster = Monster{
		Name:  "小明",
		Age:   22,
		Score: 88.8,
		Sex:   "男",
	}
	TestStruct(monster)
}
