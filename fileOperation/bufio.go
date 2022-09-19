package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//file 可以称为file对象、file指针、file文件句柄
	file, err := os.Open("d:/readme.txt")
	if err != nil {
		fmt.Println("err=", err)
	} else {
		//输出文件，看文件里是什么，看出file就是一个指针 *File
		fmt.Printf("file=%v\n", *file)
	}
	//关闭文件
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("close file err=", err)
	//}
	//当函数退出时，要及时关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄露

	//创建一个 *Reader ,是带缓冲的
	/*
		const (
			defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到换行符就结束，即一行一行读取内容
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}
