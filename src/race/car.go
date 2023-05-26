package race

import "fmt"

// if the name of the function, struct or any variable starts with uppercase
// it means that it is public, if it starts with lowercase
// it means that it is private

// Car with public access
type Car struct {
	Brand string
	Year  int
}

/* type privateCar struct {
	brand string
	year  int
}
*/

func PrintMessage(message string) {
	fmt.Println(message)
}
