//pointer is a variable which stores memory address of another variable

///declare a pointer
package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// & gets the address of variable
	//type of pointer is *int
	//zero value of a pointer is nil
}
