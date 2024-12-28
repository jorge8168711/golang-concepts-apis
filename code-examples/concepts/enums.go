package concepts

// Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed
// number of possible values, each with a distinct name. Go doesn’t have an enum type as a
// distinct language feature, but enums are simple to implement using existing language idioms.

const (
	Axe int = iota
	Sword
	WoodenStick
	Knife
)

func getDamaage(weaponType int) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 80
	case WoodenStick:
		return 40
	case Knife:
		return 20
	default:
		panic("Invalid weapon type")
	}
}

func main() {
	println("Hello, World!", getDamaage(Axe))
}
