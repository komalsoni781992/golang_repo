package compositedatatypes

import (
	"fmt"
	"reflect"
)

// TestSlice - tests slice functionality
func TestSlice() {
	var slice = []int{10, 20, 30}
	var sliceTwo []int

	//x is assigned the zero value for a slice, which is is something we haven’t seen before: nil.
	//it is slightly different from the null that’s found in other languages.
	//In Go, nil is an identifier that represents the lack of a value for some types.
	//Like the untyped numeric constants nil has no type, so it can be assigned or compared against values of different types. A nil slice contains nothing.
	//A slice is the first type that we’ve seen which isn’t comparable. It is a compiletime error to use == to see if two slices are identical or != to see if they are different.
	//The only thing you can compare a slice with is nil.

	var sparseSlice = []int{1, 5: 4, 6, 10: 100, 15}
	var matrixSlice [][]int
	slice[0] = 100

	//slice[3] = 10 - causes panic

	// len can accept alice, array, string, map, channels

	//capacity - number of consecutive memory locations reserved
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(sliceTwo, len(sliceTwo), cap(sliceTwo), sliceTwo == nil)
	fmt.Printf("%#v\n", sliceTwo)
	fmt.Println(sparseSlice, len(sparseSlice), cap(sparseSlice))
	fmt.Println(matrixSlice, len(matrixSlice), cap(matrixSlice))

	//The reflect package contains a function called DeepEqual that can compare almost anything, including slices
	fmt.Println(reflect.DeepEqual(sliceTwo, matrixSlice))

	//If you try to add additional values when the length equals the capacity, the append function uses the Go runtime(make function) to allocate a new slice with a larger capacity.
	//The values in the original slice are copied to the new slice, the new values are added to the end, and the new slice is returned
	//double the size of the slice when the capacity is less than 1024, and then grow by at least 25% afterwards
	slice = append(slice, 40)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 50, 60, 70, 80)
	fmt.Println(slice, len(slice), cap(slice))

	// appending slice with variadic operator
	slice = append(slice, sparseSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	//compile-time error if you forget to assign the value returned from append
	// Go is call by value language - Every time you pass a parameter to a function, Go makes a copy of the value that’s passed in

	//slicing
	sliceThree := sparseSlice[1:6] // index:len // this line creates a new slice
	fmt.Println(sliceThree)
	//sliceThree = sparseSlice[:]  // take the whole slice
	//sliceThree = sparseSlice[:3] // start from 0 index till the length provided
	//sliceThree = sparseSlice[2:] // from the 2nd index till the end
	fmt.Println(sliceThree)
}

// Inspect - inspects underlying array for length and capacity and memory address
func Inspect(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}
