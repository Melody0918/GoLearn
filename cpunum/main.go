package main

import (
	"fmt"
	"runtime"
)

func main() {
	//查看本地机器的逻辑CPU个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu，go1.8后默认让程序运行在多个核上，可以不用设置了
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
