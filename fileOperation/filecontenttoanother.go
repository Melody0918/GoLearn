package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将d:/abc.txt 文件内容导入到 e:/test.txt文件

	//1.首先将 d:/abc.txt 内容读取到内存
	//2.将读取到的内容写入 e:/test.txt 中
	file1path := "d:/abc.txt"
	file2path := "e:/test.txt"
	content, readErr := ioutil.ReadFile(file1path) //[]byte
	if readErr != nil {
		fmt.Println("read file err=", readErr)
		return
	}
	writeErr := ioutil.WriteFile(file2path, content, 0666) //windows下FileMode无效，可以写任意值
	if writeErr != nil {
		fmt.Println("write file err=", writeErr)
		return
	}
}
