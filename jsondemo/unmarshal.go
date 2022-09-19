package main

import (
	"encoding/json"
	"fmt"
)

type Monster1 struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func unMarshalStruct() {
	//str 在项目开发中是通过网络传输获取到的..或者是读取文件获取到的
	str := " {\"Name\":\"张三\",\"Age\":20,\"Birthday\":\"2002-1-1\",\"Sal\":10000,\"Skill\":\"Coding\"}"

	var monster Monster1

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
	}
	fmt.Println("反序列化后 monster=", monster)
}

//演示将json字符串反序列化成map
func unMarshalMap() {
	str := "{\"address\":\"湖北省\",\"age\":24,\"name\":\"李四\"}"
	var a map[string]interface{} //定义一个map

	//反序列化，不需要make，因为make操作被封装到Unmarshal函数中，反序列化底层会自动make
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
	}
	fmt.Println("反序列化后 monster=", a)
}

//演示将json字符串反序列化为切片
func unMarshalSlice() {
	str := " [{\"address\":\"北京\",\"age\":34,\"name\":\"jack\"},{\"address\":[\"上海\",\"武汉\"],\"age\":20,\"name\":\"tom\"}]"
	var slice []map[string]interface{}

	//反序列化，不需要make，因为make操作被封装到Unmarshal函数中，反序列化底层会自动make
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
	}
	fmt.Println("反序列化后 monster=", slice)
}
func main() {
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice()
}
