package basicdatatypes

import "fmt"

// TestIntegers tests Int data type for Go
func TestIntegers() {
	// int can be int, int8, int16, int32, int64, uint8 == byte , uint16, uint32 == rune, uint64, uint
	integer := 999
	//int can be _ separated for readability
	integerSeparated := 1_2_3_4
	integerSeparatedTwo := 123_456_23
	//binary - 0b, 0B, Octal - 0o, 0O, 0(Not recommended), Hexadecimal- 0x, 0X
	binaryInteger := 0b101
	binaryIntegerTwo := 0b101
	octalInteger := 0o101
	octalIntegerTwo := 0o101
	octalIntegerThree := 0101
	hexadecimalInteger := 0x101
	hexadecimalIntegerTwo := 0x101
	var a int8 = 100
	var b int16 = 100
	var c int32 = 100
	var d int64 = 100
	var e uint8 = 100
	var f uint16 = 100
	var g uint32 = 100
	var h uint64 = 100
	var i byte = 100
	var j int = 100
	// i = e is valid
	//A byte is an alias for uint8; it is legal to assign, compare, or perform mathematical operations between a byte and a uint8

	// The second special name is int. On a 32-bit CPU, int is a 32 bit signed integer like an int32. On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.

	// Because int isn’t consistent from platform to platform, it is a compile-time error to assign, compare, or perform mathematical operations between an int and an int32 or int64
	// without a type conversion
	//Integer literals default to being of int type.
	fmt.Println(integer)
	fmt.Println(integerSeparated)
	fmt.Println(integerSeparatedTwo)
	fmt.Println(binaryInteger)
	fmt.Println(binaryIntegerTwo)
	fmt.Println(octalInteger)
	fmt.Println(octalIntegerTwo)
	fmt.Println(octalIntegerThree)
	fmt.Println(hexadecimalInteger)
	fmt.Println(hexadecimalIntegerTwo)
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	// Operators available - +, -, *, /, %, +=, -=, *=, /=, %=, ==, !=, >, >=, <, <=, << , >>, &, |, ^, &^, &=, |=, ^=, &^=, <<= , >>=

	// and not &^ operatror is new
	// There are also uint and uintptr
}
