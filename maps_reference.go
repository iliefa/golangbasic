package main

import "fmt"

func main() {

	personSalary := map[string]int{
		"andreea": 1000,
		"michael": 500,
	}

	fmt.Println("original person salary ", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 3000
	fmt.Println("new person salary ", personSalary)

	//same case in functions -any change made to the map inside the function,
	//it will be visible to the caller
}
