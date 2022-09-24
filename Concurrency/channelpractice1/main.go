package main

import (
	"fmt"
	"math/rand"
	"time"
)

//完成goroutine和channel协同工作的案例，具体要求：
//1.开启一个writeData协程，向管道intChan中写入50个整数
//2.开启一个readData协程，从管道intChan中读取writeData写入的数据
//3.注意的是writeDat和readData操作的是同一个管道
//4.主线程需要等待writeData和readData协程都完成工作时才能退出

var (
	intChan  chan int  = make(chan int, 100)
	exitChan chan bool = make(chan bool, 1)
)

func writeData(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		channel <- rand.Intn(1000) + 1
	}
	close(intChan)
}

func readData(channel chan int, exitChan chan bool) {
	for {
		v, ok := <-channel
		if !ok {
			break
		}
		fmt.Println("readData读到数据:", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	//向管道写入50个整数
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		data, ok := <-exitChan
		fmt.Println("data=", data)
		fmt.Println("ok=", ok)
		if ok {
			break
		}
	}

}
