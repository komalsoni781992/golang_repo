package basicdatatypes

import "fmt"

const xValue int64 = 10 // const var-type-val
const (
	idKey   = "id"   // const var-val
	nameKey = "name" // const var-val
)
const zValue = 20 * 10 // const var-val

// ConstTest tests const values and block
func ConstTest() {
	/*
		numeric literals
		true and false
		strings
		runes
		the built-in functions complex, real, imag, len, and cap
		expressions that consist of operators and the above values
		iota - 0
	*/
	/* the Go compiler allows you to create unread constants with const. This is because constants in Go are calculated at compile time and cannot have any side-effects.
	This makes them easy to eliminate; if a constant isn’t used, it is simply not included in the compiled binary.
	*/
	const y = "hello" // const var-val
	fmt.Println(xValue)
	fmt.Println(zValue)
	fmt.Println(idKey)
	fmt.Println(nameKey)
	fmt.Println(y)

	/*
		untyped const
		const x = 10
		valid operations
		var y int = x
		var z float64 = x
		var d byte = x

		typed const
		const typedX int = 10
	*/
	var p float64 = zValue
	var q float64 = float64(xValue)
	fmt.Println(p, q)
}
