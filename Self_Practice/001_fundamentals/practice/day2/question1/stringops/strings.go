package stringops

import (
	"fmt"
)

// ReverseAndUppercase - reverses first string, uppercases second string and returns them concatenated
func ReverseAndUppercase(s1, s2 string) string {
	fmt.Println(reverseString(s1))
	fmt.Println(toUpperCase(s2))
	return reverseString(s1) + toUpperCase(s2)
}
