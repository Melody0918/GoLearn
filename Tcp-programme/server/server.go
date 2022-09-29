package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里我们循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送],那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端 %v 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return
		}
		//3.去除换行符后显示客户端发送的内容到服务器的终端,若客户端发送exit则退出
		str := string(buf[:n])
		//str = strings.TrimSpace(str)
		if str == "exit" {
			fmt.Printf("服务端监听客户端%v的数据结束！\n", conn.RemoteAddr().String())
			return
		}
		fmt.Println(str)
	}
}
func main() {
	fmt.Println("服务器开始监听....")
	//net.Listen("tcp","0.0.0.0:8888")
	//1.tcp 表示使用网络协议是tcp
	//2.0.0.0.0:8888 表示在本地监听 8888 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭

	//循环等待客户端来连接我
	for {
		//等待客户端连接
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() success conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备起一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Println("listen=", listen)
}
