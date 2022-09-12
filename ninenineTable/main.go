package main

import "fmt"

func main() {
	var i int = 1
	//var j int = 1
	for i < 10 {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
		i++
	}
}
