package main

import "fmt"

type Person struct {
	string
	int
}

//by default the name of the anonymous field is the name of it's type
func main() {
	var p1 Person
	p1.string = "zidane"
	p1.int = 32
	fmt.Println(p1)

	p2 := Person{"thiery", 45}
	fmt.Println(p2)

	//also nested structs exists
}
