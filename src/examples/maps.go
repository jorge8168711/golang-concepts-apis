package examples

import "fmt"

func Maps() {
	// the maps by default implement concurrency
	m := make(map[string]int)
	m["Juan"] = 8
	m["Jorge"] = 12

	fmt.Println(m)

	// iterate over the map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// find value
	// if the value is not found return the zero value
	// in this scenario the zero value is 0 because the dictionary value type is int
	// and the zero value for int is 0
	// when we access to the value we have to check if the value is found
	// we can use the ok variable to check if the value is found
	value, ok := m["Juan"]
	fmt.Println(ok, value)
}
