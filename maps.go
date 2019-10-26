package main

import "fmt"

//personSalary := make(map[string]int)
//creates map personSalary ,with string keys and int values

//zero value of map is nil .it has to be initialized using make,otherwise run time panic will occur

func main() {
	var personWage map[string]int
	if personWage == nil {
		fmt.Println("map is nil going to make one")
		personWage = make(map[string]int)
	}
	personWage["john"] = 1000
	personWage["andrew"] = 5000
	fmt.Println("salaries map", personWage)

	//declare during initialization
	personSalary := map[string]int{
		"eric":    100,
		"andreea": 200,
	}
	fmt.Println("salaries map", personSalary)
	employee := "eric"
	fmt.Println("salary of", employee, "is", personSalary[employee])

	// if key is not present, it will return 0
	employee = "erica"
	fmt.Println("salary of", employee, "is", personSalary[employee])

	//value, ok to see if key exists in the map
	newemployee := "george"
	value, ok := personSalary[newemployee]
	if ok == true {
		fmt.Println("salary of ", newemployee, " is", value)
	} else {
		fmt.Println(newemployee, "not found")
	}

}
