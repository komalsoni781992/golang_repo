package compositedatatypes

import "fmt"

// https://github.com/Diwakarsingh052/genspark/blob/main/8-struct/4-struct-pointers.go

type user struct {
	name string
	age  int
}

// TestPointerToSTruct -Tests updation of struct via struct pointer
func TestPointerToSTruct() {
	u := user{"jack", 18}
	fmt.Println(u)
	updateUser(&u) // passing the reference of the u object
	fmt.Println(u)
}

// passing a struct to a function which accepts struct value as a pointer
func updateUser(u *user) {
	u.age = 20 // we don not have to dereference the struct object to access the field
}
