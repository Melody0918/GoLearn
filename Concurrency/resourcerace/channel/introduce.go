package main

import "fmt"

func main() {
	//介绍+演示
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	//当我们给管道写入数据时，不能超过其容量
	//但是取出后可以继续放入
	<-intChan // 相当于丢弃数据
	intChan <- 123
	//否则会发生死锁 deadlock

	//4.看看管道的长度和cap容量
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) //2,3

	//5.从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//6.<-exitChan取空值是否会卡住,显示：fatal error: all goroutines are asleep - deadlock!
	var exitChan chan int = make(chan int, 8)
	exitChan <- 12
	exitChan <- 4
	for i := 0; i < 8; i++ {
		<-exitChan //fatal error: all goroutines are asleep - deadlock!
		fmt.Println(i)
	}

	//7.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

	num3 := <-intChan
	num4 := <-intChan
	fmt.Printf("num3=%v num4=%v\n", num3, num4) //deadlock
}
