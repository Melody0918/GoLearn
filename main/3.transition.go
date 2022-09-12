package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.423
	var b bool = true
	var myChar byte = 'h'
	var str string //空str

	// 基础数据类型转string
	// 使用第一种方法来转换 fmt.Sprintf方法

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str, str)

	fmt.Println(myChar)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q\n", str, str)

	//第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 34.5678
	var b2 = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n", str, str)

	//string转换成基本数据类型
	var str1 string = "true"
	var b1 bool
	//var e error
	// 说明：
	// 1.strconv.ParseBool(str)函数会返回两个值(value bool, err error)
	// 2.由于只想获取value bool,因此使用_忽略第二个值
	//b1, e = strconv.ParseBool(str1)
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type %T b1 = %t\n", b1, b1)
	//fmt.Printf("e type %T e = %v\n", e, e)
	var str2 string = "-1234"
	var n1 int64
	var n2 uint64

	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2, _ = strconv.ParseUint(str2, 10, 64)
	fmt.Printf("n1 type %T n1 = %d\n", n1, n1)
	fmt.Printf("n2 type %T n2 = %d\n", n2, n2)

	var str3 string = "12345.783476"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1 = %f\n", f1, f1)

	if age := 20; age > 18 {
		fmt.Println("你年纪大于18，要对自己的行为负责")
	}

	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层,输出ok1与ok2
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("未匹配到...")
	}
}
