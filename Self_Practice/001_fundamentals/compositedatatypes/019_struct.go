package compositedatatypes

import "fmt"

// You can define a struct type inside or outside of a function.
// A struct type that’s defined within a function can only be used within the function.

// TestStruct - tests struct functionality
func TestStruct() {
	type person struct {
		name string
		age  int
		pet  string
	}
	//Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value at all
	var fred person
	bob := person{}
	fmt.Printf("%#v %T %T %T\n", fred, fred.name, fred.age, fred.pet) // access of fields
	fmt.Printf("%#v\n", bob)

	//When using this struct literal format, a value for every field in the struct must be specified, and the values are assigned to the fields in the order they were
	//declared in the struct definition.

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	//When you use this style, you can leave out keys and specify the fields in any order.
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)
	//e. You cannot mix the two struct literal styles; either all of the fields are specified with keys, or none of them are
	//If you initialize a struct without using the field names and a future version of the struct adds additional fields, your code will no longer compile.
}

// TestAnonymousStruct - Introduction to anonymous struct
func TestAnonymousStruct() {
	//There are two common situations where anonymous structs are handy.
	//The first is when you translate external data into a struct or a struct into external data (like JSON or protocol buffers).
	//This is called marshaling and unmarshaling data.

	// Also in testing - slice of anonymous structs when writing table-driven tests
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}

//Whether or not a struct is comparable depends on the struct’s fields.
//Structs that are entirely composed out of comparable types are comparable; those with slice or map or channel or function fields are not
//Go does allow you to perform a type conversion from one struct type to another if the fields of both structs have the same names, order, and types

/*func StructConversion() {
	type firstPerson struct {
		name string
		age  int
	}
	type secondPerson struct {
		name string
		age  int
	}
	type thirdPerson struct {
		age  int
		name string
	}
	type fourthPerson struct {
		firstName string
		age       int
	}
	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}
	f := firstPerson{
		age:  30,
		name: "Soni",
	}
	s := secondPerson{
		age:  30,
		name: "Komal",
	}
	//f = s
}*/

// TestTypeConversionInAnonymousStruct - test anonymous structs for type conversion and comarison
func TestTypeConversionInAnonymousStruct() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	g := struct {
		name string
		age  int
	}{
		name: "Fido",
		age:  20,
	}
	fmt.Println(g)
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)

	// Anonyhmous struct initialization
	var h struct {
		name string
		age  int
	} = struct {
		name string
		age  int
	}{name: "Komal", age: 20}
	fmt.Println(h)
}
