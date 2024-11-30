package controlstructures

import (
	"fmt"
	"math/rand"
)

// TestSwitchFunctionality - switch functionality in GO
func TestSwitchFunctionality() {
	// fallthrough should be avoided
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:",
				wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// break in switch breaks case not the enclosing for loop
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			//break
			//fmt.Println("exit the loop agin!")
		default:
			fmt.Println(i, "is boring")
		}
	}

	// fix
	// break in switch breaks case not the enclosing for loop
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
			//fmt.Println("exit the loop agin!")
		default:
			fmt.Println(i, "is boring")
		}
	}

	words = []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	// clumsy
	switch a := 10; {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	// improved
	a := 10
	switch a {
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("That's too low")
	case n > 5:
		fmt.Println("That's too big:", n)
	default:
		fmt.Println("That's a good number:", n)
	}
	//Favor blank switch statements over if/else chains when you have multiple related cases.
	//Using a switch makes the comparisons more visible and reinforces that they are a related set of concerns.
}
