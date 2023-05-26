package main

import (
	"fmt"
	"golang-basics/src/examples"
)

func main() {
	// CONSTANTS
	// const pi float64 = 3.14
	// const pi2 = 3.14

	// INTS
	/* 	base := 12 // define the variable assign the value and type
	   	var height int = 14
	   	var area int
	*/

	// ZERO VALUES
	/* 	var a int     // by default is -> 0
	   	var b float64 // by default is -> 0
	   	var c string  // by default is -> ''
	   	var d bool    // by default is -> false
	*/

	/* 	fmt.Println(examples.IsPalindrome("casa"))
	   	examples.Operators()
	   	examples.Arrays()
	   	examples.IterateSlice()
	   	examples.Conditional()
	   	examples.Counter()
	   	examples.Iterate()
	   	examples.SwitchTest()
	   	examples.Structs()
	   	examples.Maps()
	   	examples.Pointers()
	   	examples.Interfaces()
	*/

	mySquare := examples.Square{Base: 2}
	myRectangle := examples.Rectangle{Base: 2, Height: 4}
	examples.CalculateArea(mySquare)
	examples.CalculateArea(myRectangle)

	// defer -> execute a function at the end of the function execution
	// is like clean up function
	// if several defer are defined, they are executed in reverse order
	// the calls are treated as a stack of calls and executed in LIFO order
	// en example is to close a file at the end of the function or  close the connection to the database
	defer fmt.Println("This is executed at the end")
}
