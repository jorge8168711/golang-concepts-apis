package main

import (
	"fmt"
)

type Color int

// this is a fmt.Stringer interface implementation
func (c Color) String() string {
	return fmt.Sprintf("Color %d", c)
}

const (
	ColorBlue  Color = iota // 0
	ColorRed                // 1
	ColorBlack              // 2
)

type Player struct {
	name        string
	health      int
	attackPower int
}

func main() {
	player1 := Player{
		name:        "Gerald",
		health:      100,
		attackPower: 20,
	}

	// print field and values of the struct
	fmt.Printf("this is the player: %+v\n", player1)

	// the println will print the string representation of the Color type
	// will look first for the String() method in the Color type
	fmt.Println(ColorBlue)
}
