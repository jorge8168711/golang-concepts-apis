package concepts

import (
	"fmt"
)

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

	// The builtin delete removes key/value pairs from a map.
	delete(m, "juan")

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions for maps.
	//n2 := map[string]int{"foo": 1, "bar": 2}
	//if maps.Equal(n, n2) {
	//	fmt.Println("n == n2")
	//}
}

func printColors(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}

func run() {
	// var otherColors map[string]string
	otherColors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
	}

	// add new key-value pair
	otherColors["white"] = "#ffffff"
	otherColors["black"] = "#000000"

	//. aceesing a key-value pair
	c, ok := otherColors["white"]
	if !ok {
		fmt.Println("color not found")
	} else {
		fmt.Println("color is", c)
	}

	// delete key-value pair
	delete(otherColors, "white")

	fmt.Println(colors)
	fmt.Println(otherColors, otherColors["white"])
	fmt.Println("HERE", colors["SOME"])

	printColors(colors)
}
