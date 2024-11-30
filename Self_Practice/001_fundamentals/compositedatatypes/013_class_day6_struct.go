package compositedatatypes

import "fmt"

// StructIntroduction - https://github.com/Diwakarsingh052/genspark/blob/main/8-struct/1-struct.go
func StructIntroduction() {
	// type typeName struct
	// group related types together in one namespace
	type Person struct {
		name    string // fields of struct
		age     int
		address string
	}

	var p Person // by default everything is intialized with its zero values
	//fmt.Printf("Msg : %#v\n", p) // print field value pairs with some extra info about the type
	fmt.Printf("Msg : %+v\n", p) // print field value pairs
	fmt.Println(p.name, ", ", p.age, ", ", p.address)

}
