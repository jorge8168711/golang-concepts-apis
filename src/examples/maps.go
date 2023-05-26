package examples

import "fmt"

// Maps are Go’s built-in associative data
// type (sometimes called hashes or dicts in other languages).

func Maps() {
	// the maps by default implement concurrency
	// Create an empty map using the builtin make:
	// make(map[key-type]val-type).

	m := make(map[string]int)

	// set key/value pairs using typical name[key] = val syntax.
	m["Juan"] = 8
	m["Jorge"] = 12

	fmt.Println(m)

	// get value for a key with name[key]
	v1 := m["Juan"]
	fmt.Println("Juan:", v1)

	// find value
	// If the key doesn’t exist, the zero value of the value type is returned.
	// in this scenario the zero value is 0 because the dictionary value type is int
	// and the zero value for int is 0.
	// When we access to the value we have to check if the value is found
	// we can use the ok variable to check if the value is found:
	value, ok := m["Juan"]
	fmt.Println(ok, value)

	// iterate over the map
	for i, v := range m {
		fmt.Println(i, v)
	}
}
