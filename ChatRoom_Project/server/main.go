package main

import (
	"GoProject/ChatRoom_Project/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn则不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err=", err)
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据 pkgLen 读取消息内容
	n, err1 := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err1 != nil {
		//fmt.Println("conn.Read fail err=", err1)
		//err = errors.New("read pkg body error")
		return
	}

	//把pkgLen反序列化成 -> message.Message
	err2 := json.Unmarshal(buf[:pkgLen], &mes)
	if err2 != nil {
		fmt.Println("json.Unmarsha err=", err2)
		return
	}
	return
}

func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环读取客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return
			} else {
				fmt.Println("readpkg err=", err)
				return
			}
		}
		fmt.Println("mes=", mes)
	}
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听......")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Listen err,err=", err)
		return
	}
	defer listen.Close()
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("listen.Accept err,err=", err1)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
