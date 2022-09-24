package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello,world")
		time.Sleep(time.Second)
	}
}
func test() {
	//这里我们可以使用defer+recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error panic: assignment to entry in nil map
}

func main() {
	//goroutine中使用recover，解决协程中出现的panic，导致程序崩溃的问题
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main()=", i)
		time.Sleep(time.Second)
	}
}
