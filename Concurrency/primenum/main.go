package main

import (
	"fmt"
	"math"
	"time"
)

func putNum(putChan chan int, exitChan chan int) {
	for i := 1; i <= 80000; i++ {
		putChan <- i
	}
	exitChan <- 1
}

func primeNum(putChan chan int, primeChan chan int, exitChan chan int) {
	for {
		v, ok := <-putChan
		var isPrime bool = true
		if !ok {
			break
		}
		//求素数
		if v == 2 {
			primeChan <- v
			continue
		}
		for i := 2; i <= int(math.Sqrt(float64(v)))+1; i++ {
			if v == 1 || v%i == 0 { //说明不是素数
				isPrime = false
				continue
			}
		}
		if isPrime {
			primeChan <- v
		}
	}
	exitChan <- 1
}

func main() {
	var intChan = make(chan int, 80000)
	var primeChan = make(chan int, 20000)
	//标识退出的管道
	var exitChan = make(chan int, 4)
	start := time.Now().Nanosecond()
	go putNum(intChan, exitChan)
	for i := 0; i < 1; i++ {
		<-exitChan
	}
	close(intChan)
	//优化前
	//go primeNum(intChan, primeChan, exitChan)
	//优化后
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//for i := 0; i < 4; i++ {
	//	<-exitChan
	//}
	//close(primeChan)
	//用匿名函数简化
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	//结束时间
	end := time.Now().Nanosecond()
	for {
		v, ok := <-primeChan
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	fmt.Println("执行时间:", end-start, "ns")
	//优化前：5069200 ns;优化后：2612100 ns;匿名函数优化后: 540500 ns
}
