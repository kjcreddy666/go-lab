package main

import "fmt"

func printFloydsTriangle(n int) {
	number := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", number)
			number++
		}
		fmt.Println()
	}
}

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)
	printFloydsTriangle(rows)
}