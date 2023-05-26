package examples

import "fmt"

func Arrays() {
	// arrays
	// on the arrays is not possible to change the length
	// only the values on the array could be changed
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// len return the length of the array
	fmt.Printf("Length %d\n", len(array))

	// cap return the max length of the array
	fmt.Printf("Max Length %d\n", cap(array))

	// use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose types to
	// build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
