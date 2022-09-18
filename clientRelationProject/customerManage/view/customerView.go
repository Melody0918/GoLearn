package main

import (
	"GoProject/clientRelationProject/customerManage/service"
	"fmt"
)

type customerView struct {
	//定义必要的字段
	key  string //接收用户输入
	loop bool   //表示是否循环显示菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("--------------客户信息管理软件----------------")
		fmt.Println("		1.添加客户")
		fmt.Println("		2.修改客户")
		fmt.Println("		3.删除客户")
		fmt.Println("		4.客户列表")
		fmt.Println("		5.退	出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.customerService.AddNewCustomer()
			fmt.Println("新增用户成功！")
		case "2":
			this.customerService.UpdateCustomer()
		case "3":
			this.customerService.DeleteCustomer()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("输入错误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你已退出客户关系管理系统...")
}

//显示所有客户的信息
func (this *customerView) list() {
	//首先，获取到当前所有的客户信息
	customers := this.customerService.List()
	//显示
	fmt.Println("--------------客户列表----------------")
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------客户列表完成----------------")
}

func main() {
	//在main函数中创建一个customerView,并运行显示主菜单..
	//这里完成对customerView结构体的customerService字段的初始化
	newcustomerservice := service.NewCustomerService()
	customerview := customerView{
		"",
		true,
		newcustomerservice,
	}
	//newcustomerservice := service.NewCustomerService()
	customerview.mainMenu()
}
