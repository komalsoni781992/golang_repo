package compositedatatypes

import "fmt"

//https://github.com/Diwakarsingh052/genspark/blob/main/8-struct/3-methods.go

// Student type with name and age unexported fields
type Student struct {
	name string
	age  int
}

//func (receiver) funcSignature {//body}
// methods could be only called using struct variable

// SayHello - method over Student receiver
func (s Student) SayHello() {
	fmt.Println("Hello, my name is", s.name)
}

// TestStructMethods - tests methods in struct
func TestStructMethods() {
	s := Student{"Jack", 20}
	s.SayHello()

}
