package concepts

import (
	"fmt"
)

type Figures2D interface {
	Area() float64
}

type Square struct {
	Base float64
}

func (s Square) Area() float64 {
	return s.Base * s.Base
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

func calculateArea(figure Figures2D) {
	fmt.Println("Area:", figure.Area())
}

func Interfaces() {
	// interface with multiple types
	myInterface := []interface{}{"Hello", 12, 4.90}
	fmt.Println(myInterface...)

	// implementing an interface
	mySquare := Square{Base: 2}
	myRectangle := Rectangle{Base: 2, Height: 4}
	calculateArea(mySquare)
	calculateArea(myRectangle)
}

// All the interfaces names ends with the suffix "er" or "able" or "ible" or "or" or "er
// An interface is simply a contract based on methods that are implemented
// and anything that implements those as defined in the interface implicitly implement said interface.
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoNumberStore struct {
	// some fields
}

func (m MongoNumberStore) GetAll() ([]int, error) {
	// logic to get all items
	return []int{1, 2, 3}, nil
}

func (m MongoNumberStore) Put(n int) error {
	// logic to put an item
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

type Position struct {
	postx int
	posy  int
}

type Entity struct {
	id   string
	name string
	Position
}

// with this approach we embed the Entity struct into the EspecialEntity struct
// is like a composition
type EspecialEntity struct {
	Entity
	someField float32
}

func InterfacesV2() {
	api := ApiServer{numberStore: MongoNumberStore{}}
	fmt.Println(api)

	e := EspecialEntity{Entity: Entity{
		id:   "1",
		name: "Jorge",
		Position: Position{
			postx: 1,
			posy:  2,
		},
	}}

	e.name = "Jorge Barron"
	e.id = "2"
	fmt.Println(e)
}
