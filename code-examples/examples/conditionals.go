package examples

import (
	"fmt"
	"log"
	"strconv"
)

func Conditional() {
	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("Value is 1")
	} else {
		fmt.Println("Value is not 1")
	}

	if value1 == 1 && value2 == 2 {
		fmt.Println("Value is 1 and value2 is 2")
	}

	if value1 == 0 || value2 == 2 {
		fmt.Println("Value is 0 or value2 is 2")
	}

	// convert txt to int
	value, err := strconv.Atoi("45")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
