package main

import "fmt"

//user defined data

//named structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

//anonymous structure
var employee struct {
	firstName string
	age       int
}

func main() {
	emp1 := Employee{
		firstName: "george",
		age:       24,
		salary:    1500,
		lastName:  "boston",
	}

	emp2 := Employee{"andreea", "smith", 29, 300}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	emp3 := Employee{"louis", "anderson", 33, 2500}
	fmt.Println("firstname", emp3.firstName)
	fmt.Println("lastname ", emp3.lastName)
	var emp4 Employee //initialized with 0
	emp4.firstName = "Jack"
	emp4.lastName = "Smith"
	fmt.Println("employee4 ", emp4)
	//create pointers to a struct
	emp5 := &Employee{"sam", "anderson", 50, 4000}
	fmt.Println("first name", (*emp5).firstName)
	fmt.Println("age", (*emp5).age)
}
