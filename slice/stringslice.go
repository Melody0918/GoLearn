package main

import "fmt"

//如果需要修改字符串，可以先将string --> []byte / 或者 []rune --> 修改 --> 重写转成string
func main() {
	str := "hello@atguigu"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	//细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
	//[]byte字节来处理，而一个汉字是3个字节，因此就会乱码
	//解决方法是 将string 转成 []rune 即可，因为 []rune 是按字符来处理，兼容汉字

	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println(str)
}
