package main

import "fmt"

func main() {

	personSalary := map[string]int{
		"andreea": 1000,
		"michael": 500,
	}
	personSalary["louis"] = 9000
	fmt.Println("map before deletion", personSalary)

	delete(personSalary, "michael")
	fmt.Println("map after deletion", personSalary)
	fmt.Println("length of the map is ", len(personSalary))
}
