package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Println(e)

	// 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Println(e2)

	// 3
	e3 := new(Employee)
	fmt.Println(*e3)
	e3.id = 1
	e3.name = "Name"
	fmt.Println(*e3)

	e4 := NewEmployee(1, "Names", false)
	fmt.Println(*e4)
}
