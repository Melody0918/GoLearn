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
	defer conn.Close()
	//2.向redis写入三对hash类型数据 [key-field-value]
	_, err1 := conn.Do("HMSet", "Monster1", "name", "tom", "age", 22, "skill", "sing songs")
	if err1 != nil {
		fmt.Println("HMSet Monster1 err,err=", err1)
		return
	}
	_, err2 := conn.Do("HMSet", "Monster2", "name", "lucy", "age", 20, "skill", "dancing")
	if err2 != nil {
		fmt.Println("HMSet Monster2 err,err=", err2)
		return
	}
	_, err3 := conn.Do("HMSet", "Monster3", "name", "john", "age", 23, "skill", "play basketball")
	if err3 != nil {
		fmt.Println("HMSet Monster3 err,err=", err3)
		return
	}
	//3.从redis中读取三对hash类型数据,并转为string类型
	monster1, err4 := redis.Strings(conn.Do("HMGet", "Monster1", "name", "age", "skill"))
	if err4 != nil {
		fmt.Println("monster1 get err,err=", err4)
		return
	}
	monster2, err5 := redis.Strings(conn.Do("HMGet", "Monster2", "name", "age", "skill"))
	if err5 != nil {
		fmt.Println("monster2 get err,err=", err5)
		return
	}
	monster3, err6 := redis.Strings(conn.Do("HMGet", "Monster3", "name", "age", "skill"))
	if err6 != nil {
		fmt.Println("monster3 get err,err=", err6)
		return
	}
	fmt.Println(monster1)
	fmt.Println(monster2)
	fmt.Println(monster3)
}
