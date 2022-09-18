package service

import (
	"GoProject/clientRelationProject/customerManage/model"
	"fmt"
	"strconv"
)

//该CustomerService,完成对Customer的操作，包括 增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段的后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 22, "1388584959", "182378@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户信息切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加新用户
func (this *CustomerService) AddNewCustomer() {
	var newCustomer model.Customer
	//显示
	fmt.Println("--------------新增客户----------------")
	//fmt.Print("请输入新增客户id:")
	//fmt.Scanln(&newCustomer.Id)
	this.customerNum++
	newCustomer.Id = this.customerNum //自动递增分配客户id
	fmt.Print("请输入新增客户姓名:")
	fmt.Scanln(&newCustomer.Name)
	fmt.Print("请输入新增客户性别:")
	fmt.Scanln(&newCustomer.Gender)
	fmt.Print("请输入新增客户年龄:")
	fmt.Scanln(&newCustomer.Age)
	fmt.Print("请输入新增客户电话号码:")
	fmt.Scanln(&newCustomer.Phone)
	fmt.Print("请输入新增客户邮箱:")
	fmt.Scanln(&newCustomer.Email)
	this.customers = append(this.customers, newCustomer)
}

//修改用户信息
func (this *CustomerService) UpdateCustomer() {
	fmt.Println("--------------修改用户信息----------------")
	fmt.Println("请输入要修改用户的Id：")
	var selectId string
	fmt.Scanln(&selectId)
	id, _ := strconv.Atoi(selectId) //string类型转为int类型
	for i, _ := range this.customers {
		if this.customers[i].Id == id {
			//修改
			this.customers[i].Update()
			fmt.Println("修改成功！")
			return
		}
	}
	fmt.Println("要修改用户Id输入错误,未找到!")
}

//删除用户信息
func (this *CustomerService) DeleteCustomer() {
	fmt.Println("--------------删除用户信息----------------")
	fmt.Println("请输入要删除的用户姓名：")
	var deleteName string
	j := 0
	deleteTag := false
	fmt.Scanln(&deleteName)
	for _, v := range this.customers {
		//移位法修改切片删除特定用户信息
		if v.Name != deleteName {
			this.customers[j] = v
			j++
		} else if j == len(this.customers)-1 {
			deleteTag = true
			this.customers = this.customers[:j]
		} else {
			deleteTag = true
		}
	}
	if !deleteTag {
		fmt.Println("删除失败，未找到要删除的用户姓名！")
	} else {
		fmt.Println("删除成功！")
	}
}
