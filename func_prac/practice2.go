package main

import (
	"fmt"
	"math/rand"
	"time"
)

//猜数游戏
func random_num() int {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100) + 1
	return num
}

func main() {
	guess_num := random_num()
	var num int
	var i int = 1
	for ; i < 11; i++ {
		fmt.Print("请猜数(1-100)：")
		fmt.Scanln(&num)
		if num == guess_num {
			fmt.Println("恭喜你猜对了！")
			break
		} else {
			fmt.Println("很遗憾，没猜对...")
		}
	}
	if i == 1 {
		fmt.Println("你真是个天才！")
	} else if i == 2 || i == 3 {
		fmt.Println("你很聪明，赶上我了！")
	} else if i >= 4 && i <= 9 {
		fmt.Println("一般般...")
	} else if i == 10 {
		fmt.Println("可算猜对啦。")
	} else {
		fmt.Println("说你啥点好呢。")
	}
}
