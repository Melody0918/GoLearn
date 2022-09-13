package main

import (
	"GoProject/encapsulation/practice/model"
	"fmt"
)

func main() {
	account := model.NewAccount("adminkp", "123456", 10000)
	if account != nil {
		fmt.Println("创建成功", *account)
	} else {
		fmt.Println("创建失败...")
	}
	account.SetAccountNum("kp180918")
	account.SetPwd("180918")
	account.SetBalance(1000)
	fmt.Println("账号：", account.GetAccountNum(), "密码：", account.GetPwd(), "余额：", account.GetBalance())
}
