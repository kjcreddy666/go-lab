package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	var a, b int
	fmt.Printf("Enter the two numbers: ")
	fmt.Scanf("%d %d", &a, &b)
	
	fmt.Printf("LCM of %d and %d is %d\n", a, b, lcm(a, b))
	fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd(a, b))
}