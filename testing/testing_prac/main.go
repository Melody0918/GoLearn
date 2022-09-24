package testing_prac

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//序列化Monster结构体并存入文件
func (monster *Monster) Store() bool {
	filepath := "../testprac.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err=", err)
		return false
	}
	defer file.Close()
	//写入序列化后内容
	str, marshalErr := json.Marshal(&monster)
	if marshalErr != nil {
		fmt.Println("marshalerr=", marshalErr)
		return false
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(string(str))

	writer.Flush()
	return true
}

//反序列化json格式为Monster结构体
func (this *Monster) ReStore(file *os.File) bool {
	reader := bufio.NewReader(file)
	readString, err := reader.ReadString('\n')
	if err == io.EOF {
		fmt.Println("文件读取完毕！")
	} else if err != io.EOF {
		fmt.Println("文件读取错误！")
		return false
	}
	err1 := json.Unmarshal([]byte(readString), &this)
	//fmt.Printf("%T", &monJson)
	if err1 != nil {
		fmt.Println("反序列化错误！")
		return false
	}
	return true
}
