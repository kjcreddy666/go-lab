package main

import "fmt"

func init() {
	fmt.Println("Son is initialised")
}

type Son struct {
    Name string
}

func (s *Son) Data() string {
    return "Son's name is: " + s.Name
}