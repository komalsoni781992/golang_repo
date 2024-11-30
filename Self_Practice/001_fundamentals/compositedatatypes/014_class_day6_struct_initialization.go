package compositedatatypes

import "fmt"

// StructInitialization - https://github.com/Diwakarsingh052/genspark/blob/main/8-struct/2-struct.go
func StructInitialization() {

	type person struct {
		name string
		age  int
	}

	var p person
	p.name = "ajay"
	p.age = 18

	p1 := person{"raj", 18}

	fmt.Printf("%T\n %T\n", p.name, p.age) // string
	fmt.Printf("%T\n", p1)                 // person struct

}
