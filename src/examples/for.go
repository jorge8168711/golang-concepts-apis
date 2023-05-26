package examples

import "fmt"

func Iterate() {
	// single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition/after for loop.
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}

func Counter() {
	// for while
	counter := 0
	for counter < 10 {
		fmt.Printf("Counter is %d\n", counter)
		counter++
	}

	for {
		fmt.Println("loop")
		break
	}
	// for forever
	/* 	counterForever := 0
	for {
		fmt.Printf("count %d\n", counterForever)
		counterForever++
	} */
}
