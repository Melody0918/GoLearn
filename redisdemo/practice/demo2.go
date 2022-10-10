package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//取出某个用户最近浏览的10个商品名
func getGoodsNames(conn redis.Conn) {
	names, err := redis.Strings(conn.Do("lrange", "LastGoodsNames", 0, 9))
	if err != nil {
		fmt.Println("lpop err ,err=", err)
		return
	}
	fmt.Println(names)
}

func main() {
	//记录用户浏览的商品信息，比如保存商品名
	//编写一个函数，可以取出某个用户最近浏览的10个商品名
	//提示，考虑使用list数据类型
	//通过go 向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	fmt.Println("请输入最近浏览的十个商品名：")
	for i := 0; i < 10; i++ {
		var goods string
		fmt.Scanln(&goods)
		_, err1 := conn.Do("lpush", "LastGoodsNames", goods)
		if err1 != nil {
			fmt.Println("lpush err,err=", err1)
			return
		}
	}
	getGoodsNames(conn)
}
