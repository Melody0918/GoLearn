package main

import "fmt"

func main() {
	//声明一个变量，保存接收用户输入的选项
	key := ""
	//声明一个变量，控制是否退出for循环
	loop := true
	//声明一个结构体保存支出或收入，收支金额和说明
	type logInfo struct {
		change    string
		changeNum float32
		info      string
	}
	//声明一个指针变量保存账户金额
	var acc float32
	var account *float32 = &acc

	//定义收支明细记录内容
	var details string
	//type logSlice []logInfo
	//var logslice logSlice

	//显示菜单
	for loop {
		var log *logInfo = &logInfo{
			"",
			0,
			"",
		}
		fmt.Println("--------------家庭收支记账软件----------------")
		fmt.Println("		1.收支明细")
		fmt.Println("		2.登记收入")
		fmt.Println("		3.登记支出")
		fmt.Println("		4.退	出")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("--------------当前收支明细记录----------------")
			fmt.Println("收支\t\t账户金额\t收支金额\t说	明")
			if details != "" {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细...来一笔吧！")
			}
		case "2":
			fmt.Println("--------------当前登记收入----------------")
			log.change = "收入"
			fmt.Print("请输入收入金额:")
			fmt.Scanln(&log.changeNum)
			*account += log.changeNum
			fmt.Print("请添加此收入说明:")
			fmt.Scanln(&log.info)
			fmt.Println("登记收入成功！")
			details += fmt.Sprintf("\n%v\t\t%v\t\t%v\t\t%v", log.change, *account, log.changeNum, log.info)
		case "3":
			fmt.Println("--------------当前登记支出----------------")
			log.change = "支出"
			fmt.Print("请输入支出金额:")
			fmt.Scanln(&log.changeNum)
			if log.changeNum > *account {
				fmt.Println("余额不足以支出")
				break
			}
			*account -= log.changeNum
			fmt.Print("请添加此支出说明:")
			fmt.Scanln(&log.info)
			fmt.Println("登记支出成功！")
			details += fmt.Sprintf("\n%v\t\t%v\t\t%v\t\t%v", log.change, *account, -log.changeNum, log.info)
		case "4":
			choice := ""
			fmt.Print("确认退出吗？y/n:")
			for {
				fmt.Scanln(&choice)
			label1:
				switch choice {
				case "y":
					fmt.Println("你已退出家庭记账软件的使用")
					return
				case "n":
					break
				default:
					fmt.Print("输入有误，请重新输入 y/n:")
					fmt.Scanln(&choice)
					goto label1
				}
				break
			}
		default:
			fmt.Println("请输入正确的选项！")
		}
	}
}
