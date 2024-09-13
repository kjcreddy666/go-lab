// to run use command go run week3.go father.go son.go

package main

import "fmt"

func main() {
    f := Father{Name: "Mr Frankly"}
    fmt.Println(f.Data())

	s := Son{Name: "Mr Tony"}
	fmt.Println(s.Data())
}