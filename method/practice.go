package main

import "fmt"

//1.编写结构体(MethodUtils),编程一个方法，方法不需要参数，在方法中打印一个10*8的矩阵，在main方法中调用该方法
type MethodUtils struct {
	hight int
	width int
}

func (method *MethodUtils) dayinjuxing() {
	width := method.width
	for method.hight > 0 {
		for method.width = width; method.width > 0; method.width-- {
			fmt.Print("*")
		}
		method.hight--
		fmt.Println()
	}
}

//2.根据行、列、字符打印对应行数和列数的字符，比如：行：3，列：2，字符*，则打印相应的效果
type MethodUtils2 struct {
	hight int
	width int
	sign  string
}

func (method2 *MethodUtils2) dayinjuxing2() {
	width := method2.width
	for method2.hight > 0 {
		for method2.width = width; method2.width > 0; method2.width-- {
			fmt.Print(method2.sign)
		}
		method2.hight--
		fmt.Println()
	}
}

type Calcuator struct {
	num1 float64
	num2 float64
}

//3.定义小小计算器结构体(Calcuator)，实现加减乘除四个功能：实现形式1：分四个方法完成；实现形式2：用一个方法搞定
func (calcuator *Calcuator) Cal(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.num1 + calcuator.num2
	case '-':
		res = calcuator.num1 - calcuator.num2
	case '*':
		res = calcuator.num1 * calcuator.num2
	case '/':
		res = calcuator.num1 / calcuator.num2
	default:
		fmt.Println("操作符输入有误！")
	}
	return res
}

func main() {
	var method MethodUtils
	var method2 MethodUtils2
	method.hight = 10
	method.width = 8
	method2.hight = 3
	method2.width = 2
	method2.sign = "*"
	(&method).dayinjuxing()
	(&method2).dayinjuxing2()
	var calcuator Calcuator
	calcuator.num1 = 5.6
	calcuator.num2 = 2.8
	fmt.Printf("计算结果为：%.2f", calcuator.Cal('*'))
}
