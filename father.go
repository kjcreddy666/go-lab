package main

import "fmt"

func init() {
	fmt.Println("Father is initialised")
}

type Father struct {
    Name string
}

func (f *Father) Data() string {
    return "Father's name is: " + f.Name
}