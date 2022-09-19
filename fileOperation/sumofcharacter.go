package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，用于保存统计结果
type CharCount struct {
	enCount    int // 记录英文个数
	numCount   int // 记录数字个数
	spaceCount int // 记录空格个数
	otherCount int // 记录其他字符个数
}

func main() {

	//思路：打开一个文件，创一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	filename := "d:/abc.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	//定义一个CharCount 实例
	var count CharCount
	//创建一个reader
	reader := bufio.NewReader(file)
	//开始循环读取filename的内容
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break //读到文件末尾就退出
		}
		//readByte := []rune(readString)//为了兼容中文
		//遍历 str,进行统计
		for _, v := range readString {

			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.enCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numCount++
			default:
				count.otherCount++
			}
		}
	}
	//输出统计的结果
	fmt.Printf("英文字符的个数为%v 数字的个数为%v 空格的个数为%v 其他字符的个数%v",
		count.enCount, count.numCount, count.spaceCount, count.otherCount)
}
