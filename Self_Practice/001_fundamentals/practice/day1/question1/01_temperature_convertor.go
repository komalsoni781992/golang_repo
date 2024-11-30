package main

import "fmt"

/*Q1. Declare a variable to represent temperature in Celsius.
Convert this temperature to Fahrenheit using the formula

Use fmt.Printf instead of Println
// fmt.Printf verbs could be found on top of fmt documentation // https://pkg.go.dev/fmt#hdr-Printing
Use type casting in Go*/

func main() {
	var celsius int64 = 101                                   //Declare a variable to represent temperature in Celsius.
	fmt.Println(ConvertCelsiusToFahrenheit(float64(celsius))) // By default drops unnecessary 0s
	fmt.Printf("%f\n", ConvertCelsiusToFahrenheit(float64(celsius)))
}

// ConvertCelsiusToFahrenheit converst temperature in float64 celsius to fahrenheit and returns float64 result
func ConvertCelsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}
