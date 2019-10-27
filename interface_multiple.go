package main

import "fmt"

type SalaryCalculator interface {
	calculateSalary() int
}

type LeaveCalculator interface {
	calculateLeave() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}
type Employee struct {
	name        string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) calculateSalary() int {
	return e.basicPay + e.pf
}

func (e Employee) calculateLeave() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	emp1 := Employee{
		name:        "john",
		basicPay:    5000,
		totalLeaves: 21,
		leavesTaken: 10,
	}

	var empOmp EmployeeOperations = emp1
	empOmp.calculateSalary()
	fmt.Printf("leaves left %d \n", empOmp.calculateLeave())
}
