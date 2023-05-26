package examples

import "fmt"

func Slices() {
	// Unlike arrays, slices are typed only by the elements
	// they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// slice
	// the slice is like a dynamic array
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice Length %d\n", len(slice))
	fmt.Printf("Slice Max Length %d\n", cap(slice))

	// slicing
	fmt.Println(slice[0])   // only one value
	fmt.Println(slice[:3])  // from 0 to 3
	fmt.Println(slice[2:4]) // from 2 to 4
	fmt.Println(slice[4:])  // from 2 to 4

	// append one element
	slice = append(slice, 7)
	fmt.Println(slice)

	// concat 2 arrays
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func IterateSlice() {
	slice := []string{"Hello", "World", "Go"}

	for _, value := range slice {
		fmt.Println(value)
	}

	// for with only the index
	for i := range slice {
		fmt.Println(i)
	}
}
