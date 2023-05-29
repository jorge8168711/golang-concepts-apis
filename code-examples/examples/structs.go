package examples

import (
	"fmt"
	"golang-concepts-apis/code-examples/types"
)

type Employee struct {
	id   int
	name string
}

// receiver function
// with * we can modify the values of the struct
// accessing to the memory address of the struct using the pointer
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

// Go’s structs are typed collections of fields. They’re useful
// for grouping data together to form records.
func Structs() {
	myCar := types.Car{Brand: "Ford", Year: 2022}
	fmt.Println(myCar)

	var otherCar types.Car
	otherCar.Brand = "Ferrari"
	fmt.Println(otherCar)
	types.PrintMessage("Hello Car")

	e := Employee{name: "John"}
	e.SetId(90)
	e.SetName("John Doe")
	fmt.Println(e, e.GetName())
}
