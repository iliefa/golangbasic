package main

import "fmt"

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contractor struct {
	empId    int
	basicPay int
}

func (p Permanent) calculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contractor) calculateSalary() int {
	return c.basicPay
}

//totalexpense is calculated iterating through slice SalaryCalculator

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("total expense is $%d", expense)
}

func main() {
	emp1 := Permanent{1, 1000, 20}
	emp2 := Permanent{2, 2000, 50}
	emp3 := Contractor{3, 3000}
	employees := []SalaryCalculator{emp1, emp2, emp3}
	totalExpense(employees)
}
