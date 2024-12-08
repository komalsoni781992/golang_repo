package pointers

import "fmt"

// TestPointers - tests pointers
func TestPointers() {
	var x int32 = 10 // var pointerToX *int32
	var y bool = true
	var z string = "A"
	pointerX := &x
	pointerY := &y
	pointerZ := &z
	fmt.Println(pointerX)
	fmt.Println(pointerY)
	fmt.Println(pointerZ)
	fmt.Println(&pointerX)
	fmt.Println(&pointerY)
	fmt.Println(&pointerZ)
}

// TestPointerOperators - tests address operator and indirection operator
func TestPointerOperators() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15
}

// TestPointerPanic - Nil check shild be done before dereferencing pointers
func TestPointerPanic() {
	var x *int
	fmt.Println(x == nil) // prints true
	fmt.Println(*x)       // panics
}

// pointer and variable should be of same type otherwise its a compile time error
// var x *int
// var y int64 = 20
// x = &y

// TestNewPointer- new returns pointer to variable of type initialized to its zero value
func TestNewPointer() {
	var x = new(int)      // rarely use new
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0
}

func TestPointerToStruct() {
	type foo struct {
		name string
	}
	x := &foo{}
	var y string
	z := &y
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(x.name)
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	/*p := person{
		FirstName:  "Pat",
		MiddleName: "Perry", // This line won't compile
		LastName:   "Peterson",
	}*/
	/*p := person{
		FirstName:  "Pat",
		MiddleName: &"Perry", // This line won't compile
		LastName:   "Peterson",
	}*/
	p := person{
		FirstName:  "Pat",
		MiddleName: z, // This line won't compile
		LastName:   "Peterson",
	}
	fmt.Println(p)

	// Use a helper function to turn a constant value into a pointer.
	s := person{
		FirstName:  "Pat",
		MiddleName: stringp("Hello"), // This line won't compile
		LastName:   "Peterson",
	}
	fmt.Println(s)
}

func stringp(s string) *string {
	return &s
}
