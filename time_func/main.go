package main

import (
	"fmt"
	"time"
)

func main() {
	//看看日期和时间相关函数和方法的使用
	//1.获取当前时间
	now := time.Now()
	fmt.Printf("now Type:%T\nnow:%v\n", now, now)

	fmt.Println("年=", time.Now().Year())
	fmt.Println("月=", time.Now().Month())
	fmt.Println("日=", time.Now().Day())
	fmt.Println("时=", time.Now().Hour())
	fmt.Println("分=", time.Now().Minute())
	fmt.Println("秒=", time.Now().Second())

	//格式化日期时间

	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
