package examples

import "fmt"

type figures2D interface {
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

func CalculateArea(figure figures2D) {
	fmt.Println("Area:", figure.Area())
}

func Interfaces() {
	myInterface := []interface{}{"Hello", 12, 4.90}
	fmt.Println(myInterface...)
}
