package concepts

import "fmt"

/**
What is a pointer?
- Eight byte long pointer that points to the memory address of the variable
- A pointer is a variable that stores the memory address of another variable.
- In computer science, a pointer is an object in many programming languages that stores a memory address.
*/

type Player struct {
	HP int
}

func (p *Player) takeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New HP ->", p.HP)
}

// if player is not a pointer, the function will create a copy of the player
// and the changes will not be reflected in the original player
// * -> dereference player
func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

type Database struct {
	user string
}

type ServerB struct {
	db *Database // uintprt -> 8 bytes long
}

func (s *ServerB) GetUserFromDB() string {
	if s.db == nil {
		panic("database is not initialized")
	}

	// golang is going to "dereference" the DB pointer
	// it's going to lookup the memory address of the pointer
	return s.db.user
}

func PointersV2() {
	p := Player{100}

	// & -> get the memory address of player
	TakeDamage(&p, 10)
	fmt.Println("Player HP ->", p.HP)
	p.takeDamage(20)

	// panic: runtime error: invalid memory address or nil pointer dereference
	// because the db is nill, we didn't provided the memory address
	// server := ServerB{}

	server := ServerB{
		db: &Database{user: "jorge"},
	}
	server.GetUserFromDB()
}
