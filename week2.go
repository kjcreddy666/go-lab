// Aim : To print pyramid of numbers

package main

import "fmt"

func printPyramid(n int) {
	for i := 1; i <= n; i++ {
		// Print spaces
		for j := 1; j <= n-i; j++ {
			fmt.Printf("  ")
		}

		// Print increasing numbers
		for j := i; j <= 2*i-1; j++ {
			fmt.Printf("%d ", j)
		}

		// Print decreasing numbers
		for j := 2*i - 2; j >= i; j-- {
			fmt.Printf("%d ", j)
		}

		// Move to the next line after each row
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