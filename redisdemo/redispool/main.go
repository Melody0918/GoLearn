package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当程序启动时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //最大空闲数
		Dial: func() (redis.Conn, error) { //初始化链接的代码，链接哪个ip
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("conn.Do Set err,err=", err)
		return
	}

	//取出
	name, err1 := redis.String(conn.Do("Get", "name"))
	if err1 != nil {
		fmt.Println("conn.Do Get err,err=", err1)
		return
	}

	fmt.Println("name=", name)
}
