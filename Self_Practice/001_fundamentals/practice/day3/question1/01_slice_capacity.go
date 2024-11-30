package main

import "fmt"

/*
Q1. Write a Go program that:

	Creates an empty slice with an initial capacity of 1.
	Appends integers from 1 to 1,000,000 to the slice.
	Tracks and prints the capacity change every time the slice's capacity increases.
	Prints the total number of capacity changes at the end.

	Formula:= (currentCap-lastCap) / lastCap * 100
	// Hint :- use type casting
*/
func main() {
	slice := make([]int, 0, 1)
	for i := 1; i < 1_000_000; i++ {
		previousCapacity := cap(slice)
		slice = append(slice, i)
		newCapacity := cap(slice)
		if previousCapacity != newCapacity {
			fmt.Printf("Previous capacity: %d , New capacity: %d , capacity change: %f\n", previousCapacity, newCapacity, float64(newCapacity-previousCapacity)/float64(previousCapacity)*100)
		}
	}
}
