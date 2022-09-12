package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//我们为了生成一个随机数，还需要个rand设置一个种子
	time.Now().Unix() //返回一个从1970：01：01的0时0分0秒到现在的秒数
	rand.Seed(time.Now().Unix())
	//随机生成1-100整数
	n := rand.Intn(100) + 1 //[0,100)+1
	fmt.Println(n)

	//随机生成1-100的数，直到生成了99这个数，看看一共用了几次
	//编写一个无限循环的控制，然后不停随机生成数，当生成99时，就退出这个无限循环
	var count int = 0
	for {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println(count)

}
