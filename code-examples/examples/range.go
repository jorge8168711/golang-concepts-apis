package examples

import "fmt"

// Range iterates over elements in a variety of data structures.

func Range() {
	// Here we use range to sum the numbers in a slice.
	numbers := []int{2, 3, 4}
	sum := 0
	// range on arrays and slices provides both the index and value for each entry
	// for index, num := range numbers
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on map iterates over key/value pairs
	keyValues := map[string]string{"a": "azopota", "w": "wrande"}
	for k, v := range keyValues {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over just the keys
	for key := range keyValues {
		fmt.Println("key:", key)
	}

	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune
	// and the second the rune itself.
	for i, c := range "wa" {
		fmt.Println(i, c)
	}
}
