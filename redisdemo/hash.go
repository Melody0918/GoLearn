package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	//通过go 向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//2.通过go 向redis写入数据 hash [key-field-val]
	_, err1 := conn.Do("HSet", "user1", "name", "john")
	if err1 != nil {
		fmt.Println("hset err=", err1)
		return
	}

	_, err2 := conn.Do("HSet", "user1", "age", 18)
	if err2 != nil {
		fmt.Println("hset err=", err2)
		return
	}
	defer conn.Close() //关闭...

	//3.通过go 向redis读取数据 hash [key-field-val]
	name, err3 := redis.String(conn.Do("HGet", "user1", "name"))
	if err3 != nil {
		fmt.Println("hget name err=", err3)
		return
	}

	age, err4 := redis.Int(conn.Do("HGet", "user1", "age"))
	if err4 != nil {
		fmt.Println("hget age err=", err4)
		return
	}
	//因为返回 r是 interface{}类型
	//因为 name 对应的值是string，因此我们需要转换
	fmt.Println("get name success", name)
	fmt.Println("get age success", age)
}
