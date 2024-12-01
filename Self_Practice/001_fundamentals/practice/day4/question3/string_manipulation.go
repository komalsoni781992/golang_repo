package main

import (
	"fmt"
	"strings"
)

type stringManipulatingFunction func(string) string

func main() {
	fmt.Println(stringManipulation(trimSpace, "   Hello World    "))
	fmt.Println(stringManipulation(toUpper, "Hello World"))
	fmt.Println(stringManipulation(greetMe, "Komal"))
}

func stringManipulation(f stringManipulatingFunction, input string) string {
	return f(input)
}

func trimSpace(input string) string {
	return strings.TrimSpace(input)
}

func toUpper(input string) string {
	return strings.ToUpper(input)
}

func greetMe(input string) string {
	return "Hello, " + input
}
