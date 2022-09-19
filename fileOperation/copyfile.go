package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer srcFile.Close()
	//通过srcFile,获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err1 := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("open file err=", err1)
		return
	}
	defer dstFile.Close()
	//通过dstFile,获取到 Writer
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}

func main() {
	//将d:/abc.txt 文件 拷贝到 e:/abc.txt

	//调用CopyFile 完成文件拷贝
	srcFile := "d:/abc.txt"
	dstFile := "e:/abc.txt"
	x, err := CopyFile(dstFile, srcFile)
	if err != nil {
		fmt.Println("copy file err=", err)
		return
	} else {
		fmt.Println("拷贝完成！")
	}
	fmt.Println(x)
}
