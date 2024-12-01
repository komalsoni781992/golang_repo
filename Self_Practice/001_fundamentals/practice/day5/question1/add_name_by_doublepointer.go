package main

import "fmt"

/*
q1. In main function create a slice of names.

	Add two elements in it.

	Create a function AddNames which appends new names to the slice
	Use double pointer concept to make AddNames function work
*/
func main() {
	names := make([]string, 0, 50)
	addName("Komal", &names)
	addName("Soni", &names)
	addName("Zeus", &names)
	fmt.Println(names)
}

func addName(name string, names *[]string) {
	*names = append(*names, name)
}
