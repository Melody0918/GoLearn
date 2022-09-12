package main

import "fmt"

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	n := fibo(4)
	fmt.Println(n)
}
