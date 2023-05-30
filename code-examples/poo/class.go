// POO
// interfaces
// inheritance
// composition
// classes
// structs

package poo

import "fmt"

/*
Go doesn't allow the inheritance go uses the composition.
*/
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
	EndDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "FullTimeEmployee"
}

// composition instead of inheritance
type TemporaryEmployee struct {
	Person
	Employee
	TaxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "TemporaryEmployee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func Classes() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Chuy"
	ftEmployee.age = 30
	ftEmployee.id = 1

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
