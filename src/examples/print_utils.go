package examples

import "fmt"

// Single params
func SingleParams(message string) {
	fmt.Println(message)
}

// multiple parameters with the same type
func MultiParams(a, b, int, message string) {
	fmt.Println(a, b, message)
}

func SingleReturn(a int) int {
	return a * 2
}

// return multiple values
func multipleReturn(a int) (c, d int) {
	return a, a * 2
}

func PrintUtils() {
	// Printf
	name := "Juan"
	num := 500
	// with value types
	// %s for string, %d for int, %v for any type
	fmt.Printf("%s has more than %d years\n", name, num)
	fmt.Printf("%v has more than %v years\n", name, num)

	// Sprintf
	message := fmt.Sprintf("%v has more than %v years", name, num)
	fmt.Println(message)

	// data types
	// with Printf we can print the type of the variable using %T
	x := 10
	fmt.Printf("x is of type %T\n", x)

	// value1, value2 := multipleReturn(2)
	value1, _ := multipleReturn(2)
	fmt.Println(value1)
}
