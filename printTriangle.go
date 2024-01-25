package main

import "fmt"

func printTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	printTriangle(5)
}
