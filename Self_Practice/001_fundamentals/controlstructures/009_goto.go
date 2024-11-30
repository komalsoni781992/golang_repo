package controlstructures

import (
	"fmt"
	"math/rand"
)

// TestGoto - tests goto functionality
func TestGoto() {
	/*a := 10
		goto skip
		b := 20
	skip:
		c := 30
		fmt.Println(a, b, c)
		if c > a {
			goto inner
		}
		if a < b {
		inner:
			fmt.Println("a is less than b")
		}*/
	// valid use
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

	//floatBits method in the file atof.go in the strconv package in the standard library usess goto to simplify logic
}
