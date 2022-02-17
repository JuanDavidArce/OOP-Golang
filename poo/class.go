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
}

func GetMessage(p Person) {
	fmt.Println("%s with age %d", p.name, p.age)
}
func main() {
	ftEmplyee := FullTimeEmployee{}
	ftEmplyee.name = "Juan David"
	ftEmplyee.age = 2
	ftEmplyee.id = 15
	fmt.Println(ftEmplyee)
	GetMessage(ftEmplyee)
}
