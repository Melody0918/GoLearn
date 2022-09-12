package main

import "fmt"

func test01(arr *[3]int) {
	(*arr)[0] = 88
}

func main() {
	arr := [3]int{23, 44, 62}
	test01(&arr)
	fmt.Println(arr)
	fmt.Println('A' + byte(1))

	var myChars [26]string
	for i := 0; i < 26; i++ {
		myChars[i] = string([]byte{'A' + byte(i)})
	}
	fmt.Println(myChars)
}
