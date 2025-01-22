package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Hiroki",
		phone:        "09000000000",
	}
	
	employeeList = append(employeeList, employee1)

	fmt.Println("employeeList ", employeeList)
}