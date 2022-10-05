package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {

	//通过go 向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//2.通过go 向redis写入数据 string [key-val]
	_, err1 := conn.Do("Set", "name", "tomjerry")
	if err1 != nil {
		fmt.Println("set err=", err1)
		return
	}
	_, err2 := conn.Do("Set", "job", "coding")
	if err2 != nil {
		fmt.Println("set err=", err2)
		return
	}
	defer conn.Close() //关闭...

	//3.通过go 向redis读取数据 string [key-val]
	r, err3 := redis.String(conn.Do("Get", "name"))
	if err3 != nil {
		fmt.Println("get err=", err3)
		return
	}
	//5.给数据设置有效时间
	_, err4 := conn.Do("expire", "job", 5)
	if err4 != nil {
		fmt.Println("get err=", err4)
		return
	}
	time.Sleep(5 * time.Second)
	expire, _ := redis.String(conn.Do("Get", "job"))
	fmt.Println("expire=", expire)

	//因为返回 r是 interface{}类型
	//因为 name 对应的值是string，因此我们需要转换
	fmt.Println("get name success", r)

	//4.批量Set/Get
	_, err5 := conn.Do("HMSet", "user1", "name", "tom", "age", 22)
	if err5 != nil {
		fmt.Println("HMSet err=", err5)
		return
	}

	rs, err6 := redis.Strings(conn.Do("HMGet", "user1", "name", "age"))
	if err6 != nil {
		fmt.Println("HMSet err=", err6)
		return
	}
	fmt.Println(rs)
	for i, v := range rs {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
