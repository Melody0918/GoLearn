package main

import "fmt"

//practice1作业1

var (
	numChan  = make(chan int, 2000)
	resChan  = make(chan map[int]int, 2000)
	exitChan = make(chan int, 8)
)

func storeNum(this chan int) {
	for i := 1; i <= 2000; i++ {
		this <- i
	}
	close(this)
}

//从numChan取出数并计算
func popAndCal(this chan int) {
	for {
		var sum int
		v, ok := <-this
		if !ok {
			break
		}
		for i := 1; i <= v; i++ {
			sum += i
		}
		resChan <- map[int]int{v: sum}
	}
	exitChan <- 1
}
func main() {
	go storeNum(numChan)
	//开启8个协程，来计算每个数值的累加
	for i := 0; i < 8; i++ {
		go popAndCal(numChan)
	}
	//循环请求8次，等待8次输出结束后，证明写入操作完成，这时可以继续主线程的下一步读取操作
	for i := 0; i < 8; i++ {
		<-exitChan
	}
	//关闭结果写入口
	close(resChan)
	//遍历结果集
	for {
		v, ok := <-resChan
		if ok {
			for i, j := range v {
				fmt.Printf("res[%v]=%v\n", i, j)
			}
		} else {
			break
		}
	}
}
