package examples

import "fmt"

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
func Slices() {
	// Unlike arrays, slices are typed only by the elements
	// they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Empty slice with non-zero length using the builtin make.
	s = make([]string, 5)
	fmt.Println("empty:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// slice
	// the slice is like a dynamic array
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice Length %d\n", len(slice))
	fmt.Printf("Slice Max Length %d\n", cap(slice))

	// The builtin append returns a slice containing one or more new values.
	// Note that we need to accept a return value
	// from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// Slices can also be copy’d.
	// Here we create an empty slice c
	// of the same length as s
	// and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slice operator
	// slice[low:high]
	fmt.Println(slice[0])   // only one value
	fmt.Println(slice[:3])  // from 0 to 3
	fmt.Println(slice[2:4]) // from 2 to 4
	fmt.Println(slice[4:])  // from 2 to 4

	// multi-dimensional data structures
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

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
