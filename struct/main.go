package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	//1.创建一个Monster变量
	monster := Monster{"牛魔", 500, "顶飞"}

	//2.将monster变量序列化为json格式字符串
	// json.Marshal 函数中使用反射，这个讲解反射时，会详细介绍
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr)) //{"name":"牛魔","age":500,"skill":"顶飞"}
}
