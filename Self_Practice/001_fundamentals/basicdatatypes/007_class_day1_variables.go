package basicdatatypes

import "fmt"

// go run fileName.go // to run your programs
// GO is a statically compiled language
var name string

// TestVariableInitialization - shorthand operator, shadowing, local and global variable, string, int https://github.com/Diwakarsingh052/genspark/blob/main/1-variables/1-var.go
func TestVariableInitialization() {
	var a int = 10
	var b string = "ajay"
	var c = "rahul"

	// in local functions use the below way (preferred way)
	// when you have to assign the value directly
	// shorthand operator
	d := 100 // go compiler would infer the type automatically from the right side value
	{
		// don't do it, it can cause bugs
		d := "hello" // shadowing // not recommended
		fmt.Println(d)
	}
	fmt.Println(a, b, c, d)
	fmt.Println("Hello, World!")
	fmt.Println(name)
}
