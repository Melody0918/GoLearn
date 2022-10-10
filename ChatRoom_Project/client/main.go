package main

import "fmt"

//定义两个变量，一个表示用户id，一个表示用户密码
var userId string
var password string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------------------欢迎登录多人聊天系统-------------------")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			//说明用户要登陆
			fmt.Println("请输入用户id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanln(&password)
			//先把登陆的函数写到另外一个文件，比如login.go
			err := login(userId, password)
			if err != nil {
				fmt.Println("登录失败")
				return
			} else {
				fmt.Println("登陆成功")
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误，请重新输入！")
		}
	}
}
