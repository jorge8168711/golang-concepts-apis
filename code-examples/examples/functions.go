package examples

import (
	"fmt"
	"time"
)

func Plus(a int, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go,
// for example to return both result and error
// values from a function.

// return two int values
func someValues() (int, int) {
	return 3, 7
}

func SomeFunc() {
	a, b := someValues()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier _.
	_, c := someValues()
	fmt.Println(c)
}

// Variadic functions can be called with any number of
// trailing arguments. For example, fmt.Println is a common
// variadic function.
func VariadicFunc(numbers ...int) {
	// the numbers argument is treated as a slice
	fmt.Print(numbers, " ")
	total := 0

	for _, num := range numbers {
		total += num
	}

	fmt.Println(total)

	// using a slice as parameters
	// someNumbers := []int{1, 2, 3, 4}
	// VariadicFunc(someNumbers...)
}

func someNumbers(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4

	return
}

func AnonymousFunctions() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)

	// anonymous function with a go routine
	c := make(chan int)
	go func() {
		fmt.Println("the function starts")
		time.Sleep(2 * time.Second)
		fmt.Println("the function ends")
		c <- 1
	}()

	fmt.Println(<-c)
	foo, bar, baz := someNumbers(2)
	fmt.Println(foo, bar, baz)
}
