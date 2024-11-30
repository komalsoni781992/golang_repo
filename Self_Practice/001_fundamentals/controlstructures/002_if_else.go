package controlstructures

import (
	"fmt"

	"math/rand"
)

// TestIfElse - if else in Go
func TestIfElse() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

// TestScopedIfElse - scope of n is in if and else blocks only
func TestScopedIfElse() {
	//Technically, you can put any simple statement before the comparison in an if statement.
	//This includes things like a function call that doesn’t return a value or assigning a new value to an existing variable.
	//But don’t do this. Only use this feature to define new variables that are scoped to the if/else statements; anything else would be confusing.
	//Also be aware that just like any other block, a variable declared as part of an if statement will shadow variables with the same name that are declared in containing blocks.
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	//fmt.Println(n) - compile time error
}

/*1. What does the following program print?
i := 10
if i > 10 {
 fmt.Println("Big")
} else {
 fmt.Println("Small")
}*/

// PrintSmall - prints SMall
func PrintSmall() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

// LeapYear - find if year is leap year
func LeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
