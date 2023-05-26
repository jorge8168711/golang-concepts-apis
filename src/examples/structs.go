package examples

import (
	"fmt"

	"golang-basics/src/race"
)

// Go’s structs are typed collections of fields. They’re useful
// for grouping data together to form records.
func Structs() {
	myCar := race.Car{Brand: "Ford", Year: 2022}
	fmt.Println(myCar)

	var otherCar race.Car
	otherCar.Brand = "Ferrari"
	fmt.Println(otherCar)
	race.PrintMessage("Hello Car")
}
