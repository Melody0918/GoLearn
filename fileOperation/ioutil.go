package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "d:/readme.txt"
	content, err := ioutil.ReadFile(file) //文件的打开和关闭都被隐藏到ioutil.ReadFile函数中
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v", string(content)) //[]byte
	//文件的Open和Close被封装到 ReadFile 函数内部
}
