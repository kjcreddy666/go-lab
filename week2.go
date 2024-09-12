// Aim : To print pyramid of numbers

package main

import "fmt"

func printPyramid(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Printf("  ")
		}

		for j := i; j <= 2*i-1; j++ {
			fmt.Printf("%d ", j)
		}

		for j := 2*i - 2; j >= i; j-- {
			fmt.Printf("%d ", j)
		}

		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Enter the value of n: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Pyramid of %d:\n", n)
	printPyramid(n)
}