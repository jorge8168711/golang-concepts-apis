package examples

import "fmt"

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
