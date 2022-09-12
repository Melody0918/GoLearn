package main

import "fmt"

func swap(n1 *int, n2 *int) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}

func main() {
	var x int = 34
	var y int = 78
	swap(&x, &y)
	fmt.Println("交换后x=", x, "y=", y)
}
