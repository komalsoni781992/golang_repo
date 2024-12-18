package main

import (
	"fmt"
	"strconv"
)

func main() {
	if v, err := ParseStringToFloat64("abc"); err == nil {
		fmt.Println(v)
	} else {
		panic("Invalid input")
	}
}

// ParseStringToFloat64 - parese string to float64
func ParseStringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
