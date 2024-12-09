package main

import "fmt"

/*q3. Create a new custom type based on float64 for handling temperatures in Celsius.
  Implement the Following Methods (not functions):
  Method 1: ToFahrenheit
  Description: Converts the Celsius temperature to Fahrenheit.
  Signature: ToFahrenheit() float64
  Method 2: IsFreezing
  Description: Checks if the temperature is at or below the freezing point (0°C).
  Signature: IsFreezing() bool*/

type temperature float64

func (t temperature) ToFahrenheit() float64 {
	return float64(t*9.0/5.0) + 32
}

func (t temperature) IsFreezing() bool {
	if t <= 0 {
		return true
	}
	return false
}

func main() {
	var t temperature = -2
	fmt.Println(t.ToFahrenheit())
	fmt.Println(t.IsFreezing())
}
