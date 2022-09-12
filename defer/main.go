package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈中(defer栈)
	// 当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈
	defer fmt.Println("ok1 n1=", n1) //defer 3.ok1 n1= 20
	defer fmt.Println("ok2 n2=", n2) //defer 2.ok2 n2= 30
	res := n1 + n2
	fmt.Println("ok3 res=", res) //1.ok3 res= 50
	return res

}

func main() {
	res := sum(20, 30)
	fmt.Println("res=", res) //4.res= 50
}
