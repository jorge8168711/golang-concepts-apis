package examples

import "fmt"

func SwitchTest() {
	// module := 5 % 2
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	// without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("Greater than 100")
	case value < 0:
		fmt.Println("Less than 0")
	default:
		fmt.Println("Less than 100")
	}
}
