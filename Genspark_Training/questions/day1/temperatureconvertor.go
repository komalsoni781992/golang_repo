package main

import "fmt"

/*Q1. Declare a variable to represent temperature in Celsius.
Convert this temperature to Fahrenheit using the formula

Use fmt.Printf instead of Println
// fmt.Printf verbs could be found on top of fmt documentation // https://pkg.go.dev/fmt#hdr-Printing
Use type casting in Go*/

func main() {
	var celsius int64 = 100                                      //Declare a variable to represent temperature in Celsius.
	var fahrenheit float64 = (float64(celsius) * 9.0 / 5.0) + 32 //Convert this temperature to Fahrenheit using the formula
	fmt.Printf("%f", fahrenheit)                                 //Use fmt.Printf instead of Println
}
