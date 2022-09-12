package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//统计字符串的长度，按字节 len(str)
	str := "hello呀"                   //8,golang 的编码统一为utf-8(ascii的字符(字母和数字)占一个字节，汉字占3个字节)
	fmt.Println("str len=", len(str)) //8

	// 字符串转整数：n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("1234")
	if err != nil {
		fmt.Println("转化错误", err)
	} else {
		fmt.Println("转换结果是：", n)
	}

	//整数转字符串 str = strconv.Itoa(12345)
	str1 := strconv.Itoa(12346)
	fmt.Printf("str1=%T str1=%v\n", str1, str1)

	//字符串转[]byte:var bytes = [] byte("hello")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%c\n", bytes)

	//[]byte转字符串：str = string([]byte{97,98,99})
	str = string([]byte{87, 88, 89})
	fmt.Println("str=", str)

	//查找子串是否在指定的字符串中：strings.Contains("seafood","foo") //true
	b := strings.Contains("seafood", "foo")
	fmt.Println("子串是否在指定的字符串中 ", b)

	str = strings.TrimSpace(" tn a long gopher sdj 	")
	fmt.Printf("str=%q", str)
}
