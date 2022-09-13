package model

import "fmt"

type account struct {
	accountNum string
	pwd        string
	balance    float64
}

func NewAccount(a string, p string, b float64) *account {
	if len(a) > 10 && len(a) < 6 {
		return nil
	}
	if len(p) != 6 {
		return nil
	}
	if b < 20 {
		return nil
	}

	return &account{
		accountNum: a,
		pwd:        p,
		balance:    b,
	}
}

//工厂模式的函数-构造函数
func (a *account) SetAccountNum(accnum string) {
	if len(accnum) >= 6 && len(accnum) <= 10 {
		a.accountNum = accnum
	} else {
		fmt.Println("账号长度不符合规范...")
	}
}

func (a *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		a.pwd = pwd
	} else {
		fmt.Println("密码格式有误...")
	}
}
func (a *account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("余额输入有误...")
	}
}
func (a *account) GetAccountNum() string {
	return a.accountNum
}
func (a *account) GetPwd() string {
	return a.pwd
}
func (a *account) GetBalance() float64 {
	return a.balance
}
