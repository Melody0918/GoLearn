package main

import (
	"GoProject/ChatRoom_Project/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

//写一个登陆函数
func login(userId string, password string) (err error) {
	//下一步就要开始定协议
	//fmt.Printf("userId=%s userPwd=%s\n", userId, password)
	//return nil

	//1.连接到服务器
	conn, err1 := net.Dial("tcp", "localhost:8889")
	if err1 != nil {
		fmt.Println("net.Dial err,err=", err1)
		return err1
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = password

	//4.将loginMes序列化
	data, err2 := json.Marshal(loginMes)
	if err2 != nil {
		fmt.Println("json.Marshal err,err=", err2)
		return err2
	}
	//5.将 mes进行序列化
	mes.Data = string(data)
	//6.将 mes 进行序列化
	data1, err3 := json.Marshal(mes)
	if err3 != nil {
		fmt.Println("json.Marshal(mes) err,err=", err3)
		return err3
	}
	//7.到这个时候 data1 就是我们要发送的消息
	//7.1 先把 data1 的长度发送给服务器
	// 先获取到 data 的长度->转成一个表示长度的byte切片(用binary包中的PutUint32)
	//conn.Write(len(data1))
	var pkgLen uint32
	pkgLen = uint32(len(data1))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	//发送长度
	n, err4 := conn.Write(buf[:4])
	if n != 4 || err4 != nil {
		fmt.Println("conn.Write(bytes) fail,err=", err4)
		return err4
	}
	//发送消息本身
	_, err5 := conn.Write(data1)
	if err5 != nil {
		fmt.Println("conn.Write(data1) fail err=", err5)
		return
	}
	fmt.Println("客户端，发送消息的长度=", len(data1))
	fmt.Println("客户端，发送消息的内容=", string(data1))

	//休眠20秒
	time.Sleep(10 * time.Second)
	fmt.Println("休眠了10秒")
	//这里还需要处理服务器端返回的消息
	return nil
}
