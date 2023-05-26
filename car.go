package race

import "fmt"

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
