package main

import "fmt"

func main() {
	b := 255
	a := &b //a is the address of b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is ", *a)
	*a++
	fmt.Println("new value of b is", *a)
}
