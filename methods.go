package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("salary of %s is %d%s", e.name, e.salary, e.currency)
}
func main() {
	e1 := Employee{"john", 25, " dollars"}
	e1.displaySalary()
	//changes made inside a method with a pointer receiver is visible to the caller
	//whereas this is not the case in value receiver.

}
