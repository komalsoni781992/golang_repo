package basicdatatypes

import "fmt"

// TestString tests string data type for Go
func TestString() {
	stringOne := "Greetings and\n\"Salutations\"" // normal string
	stringTwo := `Greetings and
	"Salutations"` //raw string
	fmt.Println(stringOne)
	fmt.Println(stringTwo)
	fmt.Println(len(stringOne))
	fmt.Println(len(stringTwo))        // raw string treats 4 characters as 1 tab character
	fmt.Println(stringOne + stringTwo) //concatenation
}
