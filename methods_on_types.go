package main

import "fmt"

//define methods on non-struct types
//definition of receiver type and def of method should be in the same package
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}
func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum1 := num1.add(num2)
	fmt.Println("sum is ", sum1)
}
