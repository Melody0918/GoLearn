package main

import (
	"encoding/json"
	"fmt"
)

//首先定义一个结构体
type Monster struct {
	Name     string `json:"name"` //反射机制
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	//演示
	var monster Monster = Monster{
		Name:     "张三",
		Age:      20,
		Birthday: "2002-1-1",
		Sal:      10000.0,
		Skill:    "Coding",
	}
	//将其序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	//输出序列化后的结果
	fmt.Println("data struct序列化后为", string(data))
}

//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "李四"
	a["age"] = 24
	a["address"] = "湖北省"

	//将map进行序列化
	//将其序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	//输出序列化后的结果
	fmt.Println("a map序列化后为", string(data)) //数据顺序不保证按顺序
}

//对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 34
	m1["address"] = "北京"
	slice = append(slice, m1)
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"上海", "武汉"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	//输出序列化后的结果
	fmt.Println("slice 序列化后为", string(data))
}

//对基本数据类型序列化,意义不大，就是转为string类型
func testFloaat64() {
	var num1 float64
	num1 = 3478.3784
	//将其序列化
	data, err := json.Marshal(&num1)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	//输出序列化后的结果
	fmt.Println("a map序列化后为", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloaat64()
}
