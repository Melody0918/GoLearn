package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//打开一个存在的文件，在原来的内容追加内容'ABC | ENGLISH!'
	//1.打开文件d:/abc.txt
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666) //读写模式打开文件，在文件原来的内容上追加
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//创建一个 *Reader ,是带缓冲的
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到换行符就结束，即一行一行读取内容
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容显示到终端
		fmt.Print(str)
	}
	//准备写入内容
	str := "hello,kp\n"
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
