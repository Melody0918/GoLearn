package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//创建一个新文件，写入内容 "hello,golang"
	//1.打开文件d:/abc.txt
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666) //将文件原来的内容清空
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入内容
	str := "hello,go\n"
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//writer.Write([]byte(str))
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriterString方法时，内容是先写到缓存的
	//所以需要调用Flush方法，将缓存的数据真正写到文件中，否则文件中会没有数据！
	writer.Flush()
}
