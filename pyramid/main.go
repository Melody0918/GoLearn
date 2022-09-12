package main

import "fmt"

func main() {
	var floor int = 1
	var num int = 9
	for ; floor <= num; floor++ {
		for i := 1; i <= num-floor; i++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*floor-1; j++ {
			if j == 1 || j == 2*floor-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
