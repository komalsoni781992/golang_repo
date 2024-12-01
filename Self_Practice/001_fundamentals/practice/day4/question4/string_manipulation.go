package main

import (
	"fmt"
	"strings"
)

type stringManipulatingFunction func(string) string

func main() {
	fmt.Println(stringManipulation(trimSpace(), "   Hello World    "))
	fmt.Println(stringManipulation(toUpper(), "Hello World"))
	fmt.Println(stringManipulation(greetMe(), "Komal"))
}

func stringManipulation(f stringManipulatingFunction, input string) string {
	return f(input)
}

func trimSpace() stringManipulatingFunction {
	return func(input string) string {
		return strings.TrimSpace(input)
	}
}

func toUpper() stringManipulatingFunction {
	return func(input string) string {
		return strings.ToUpper(input)
	}
}

func greetMe() stringManipulatingFunction {
	return func(input string) string {
		return "Hello, " + input
	}
}
