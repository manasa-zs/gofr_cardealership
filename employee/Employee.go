package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func EmployeeDetails(e employee) (bool, employee) {
	if e.age < 22 {
		return false, employee{}
	}
	return true, e
}

func main() {
	var (
		int [20]a
	)
	if a == b {
		fmt.Println("equal")
	}
	fmt.Println(EmployeeDetails(employee{"abc", 20, 200000}))
}
