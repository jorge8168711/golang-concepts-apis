package poo

import "fmt"

type employee struct {
	id       int
	name     string
	vacation bool
}

func (e employee) getName() string {
	return e.name
}

// with *employee we return the value of the memory address
func newEmployee(id int, name string, vacation bool) *employee {
	// we return a pointer to the struct
	// this will allow us to modify the values of the struct
	return &employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func Constructors() {
	e1 := employee{}
	e2 := employee{id: 1, name: "John", vacation: true}

	// when the new keyword is used returns a pointer to the struct
	e3 := new(employee)
	e4 := newEmployee(2, "Jane", false)

	fmt.Println(e1, e2, *e3, *e4, e4.getName())
}
