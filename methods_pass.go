package main

import "fmt"

type Employee struct {
	name string
	age  int
}

// method with value receiver

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e1 := Employee{"Mark", 50}
	fmt.Println("name before changing", e1.name)
	e1.changeName("gogu")
	fmt.Println("name after changing", e1.name)
	e1.changeAge(30)
	fmt.Println("age after changing", e1.age)
}
