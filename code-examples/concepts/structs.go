package concepts

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

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

// receiver functions
func (p person) print() {
	fmt.Printf("%+v", p)
}

// when the receiver funciton is called with a pointer type, the value is passed by reference automatically
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func structsExample() {
	// first way to declare a struct
	joan := person{"Joan", "Smith", contactInfo{"someEmail@yopmail.com", 12345}}

	// second way to declare a struct
	jorge := person{
		firstName: "Jorge",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "someemail@email.com",
			zipCode: 12345,
		},
	}

	// third way to declare a struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Smith"
	alex.contactInfo.email = "some@mail.com"
	alex.contactInfo.zipCode = 12345

	// alexPointer := &alex
	// alexPointer.updateName("Alexa")
	alex.updateName("Alexa")

	fmt.Println(joan)
	fmt.Println(jorge)

	// %+v will print the field names
	alex.print()
}
