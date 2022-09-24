package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中
//最后显示出来，要求使用goroutine完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们启动多个协程，统一将结果放入到map中
//3.map应该做成全局
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁lock
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	//这里我们开启多个协程完成这个任务[20个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	//休眠10秒钟？用互斥锁仍未解决问题
	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%v]=%v\n", i, v)
	}
	lock.Unlock()
}
