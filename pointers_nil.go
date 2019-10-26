package main

import "fmt"

func main() {

	a := 50
	var b *int
	if b == nil {
		fmt.Println("value of b is", b)
		b = &a
		fmt.Println("b after initialization", b)
	}
	// new is used to create pointers
	size := new(int)
	fmt.Printf("size value is %d, type is %T, address is %v \n", *size, size, size)
	*size = 85
	fmt.Printf("new value is ", size)
}
