package utils

import "fmt"

type FamilyAccount struct {
	//声明必须的字段
	key       string
	loop      bool
	change    string
	changeNum float32
	info      string
	account   float32
	details   string
	choice    string
}

//编写一个工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		"",
		true,
		"",
		0,
		"",
		0,
		"",
		"",
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------当前收支明细记录----------------")
	fmt.Println("收支\t\t账户金额\t收支金额\t说	明")
	if this.details != "" {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

//将登记收入写成一个方法
func (this *FamilyAccount) income() {
	fmt.Println("--------------当前登记收入----------------")
	this.change = "收入"
	fmt.Print("请输入收入金额:")
	fmt.Scanln(&this.changeNum)
	this.account += this.changeNum
	fmt.Print("请添加此收入说明:")
	fmt.Scanln(&this.info)
	fmt.Println("登记收入成功！")
	this.details += fmt.Sprintf("\n%v\t\t%v\t\t%v\t\t%v", this.change, this.account, this.changeNum, this.info)
}

//将登记支出写成一个方法
func (this *FamilyAccount) outPut() {
	fmt.Println("--------------当前登记支出----------------")
	this.change = "支出"
	fmt.Print("请输入支出金额:")
	fmt.Scanln(&this.changeNum)
	if this.changeNum > this.account {
		fmt.Println("余额不足以支出")
		return
	}
	this.account -= this.changeNum
	fmt.Print("请添加此支出说明:")
	fmt.Scanln(&this.info)
	fmt.Println("登记支出成功！")
	this.details += fmt.Sprintf("\n%v\t\t%v\t\t%v\t\t%v", this.change, this.account, -this.changeNum, this.info)
}

//将退出写成一个方法
func (this *FamilyAccount) exit() {
	fmt.Print("确认退出吗？y/n:")
	for {
		fmt.Scanln(&this.choice)
	label1:
		switch this.choice {
		case "y", "n":
			return
		default:
			fmt.Print("输入有误，请重新输入 y/n:")
			fmt.Scanln(&this.choice)
			goto label1
		}
	}
}

//给该结构体绑定方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for this.loop {
		fmt.Println("--------------家庭收支记账软件----------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.登记支出")
		fmt.Println("		4.退	出")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.outPut()
		case "4":
			this.exit()
			if this.choice == "y" {
				fmt.Println("你已退出家庭记账软件的使用")
				return
			}
		default:
			fmt.Println("请输入正确的选项！")
		}
	}
}
