package controlstructures

import "fmt"

// TestShortIf - https://github.com/Diwakarsingh052/genspark/blob/main/2-control-flow/short-if.go
func TestShortIf() {
	if i := 10; i < 10 {
		fmt.Println("i is less than 10")
	}
	//a, _ := f() // assume f() return two values, the first value would be stored in a var
	// , and the second value would be ignored because of _(underscore)

	// use short if
	// call the println func and check if it has written no values over the stdout
	// ignore the err value from the println function
	// if no values are written return otherwise print the values

	if n, _ := fmt.Println(); n == 1 {
		fmt.Println("no values written")
	} else {
		fmt.Println("values written", n)
	}

	// any variable created inside if block lives until if block is not over
	// after that, we cant access them
	fmt.Println("end of main")

}
