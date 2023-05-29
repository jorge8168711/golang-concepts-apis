package examples

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// the String() method is used to print the struct
// we can override the method String() of the struct to customize the print
func (customPc pc) String() string {
	return fmt.Sprintf("RAM: %d, DISK: %d", customPc.ram, customPc.disk)
}

// with this we can add methods to the struct
func (customPc pc) ping() {
	fmt.Println(customPc.brand, "ping")
}

// with the * we can modify the values of the struct
// accessing to the memory address of the struct using the pointer
func (customPc *pc) duplicateRAM() {
	customPc.ram = customPc.ram * 2
}

func Pointers() {
	// the pointers are used to reference a memory address
	a := 50
	// b is a pointer to access to the memory address of a
	// with the & we get the memory address of a
	b := &a

	fmt.Println(b)
	// *b return the value of the memory address
	fmt.Println(*b)

	myPC := pc{ram: 8, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
