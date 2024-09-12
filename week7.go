package main

import "fmt"

func isPalindrome(num int) bool {
	original := num
	reversed := 0

	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	return original == reversed
}

func main() {
	var num int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome")
	} else {
		fmt.Println(num, "is not a palindrome")
	}
}