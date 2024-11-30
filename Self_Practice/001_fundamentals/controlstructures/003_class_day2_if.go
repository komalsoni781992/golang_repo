package controlstructures

import "fmt"

// TestIfElseIfLadder - https://github.com/Diwakarsingh052/genspark/blob/main/2-control-flow/if.go
func TestIfElseIfLadder() {
	a := 10
	b := 20
	c := 30

	if a > b && a == c {
		fmt.Println("b is smallest and a and c are equal")
	} else if b > a && b != c {
		fmt.Println("b is larger than a")
	} else {
		fmt.Println("print anything")
	}
}
