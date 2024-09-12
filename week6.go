package main

import "fmt"

func main() {
	var str1, str2 string

	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	result := str1 + str2

	fmt.Println("Concatenated string:", result)
}