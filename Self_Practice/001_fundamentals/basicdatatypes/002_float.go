package basicdatatypes

import "fmt"

// TestFloatingPoint tests Int data type for Go
func TestFloatingPoint() {
	floatingPointTwo := 1.0
	floatingPointFour := 123_456_23.24_22
	floatingPointExponent := 6.03e23
	fmt.Println(floatingPointTwo)
	fmt.Println(floatingPointFour)
	fmt.Println(floatingPointExponent)

	//You also have the option to write them in hexadecimal by using the 0x prefix and the letter p for indicating any exponent.

	//As per IEEE 754 -  1 bit is used to represent the sign (positive or negative), 11 are used to represent a base-2 exponent,and
	//52 bits are used to represent the number in a normalized format (called the mantissa)

	// floats should not be compared rather their differnece should be compared with epsilon as they are not exact
}
