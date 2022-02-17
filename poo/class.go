package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}
func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary time employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan David"
	ftEmployee.age = 2
	ftEmployee.id = 15
	fmt.Println(ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
