package examples

import "fmt"

func Closure() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func ClosureExample() {
	counter := Closure()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
