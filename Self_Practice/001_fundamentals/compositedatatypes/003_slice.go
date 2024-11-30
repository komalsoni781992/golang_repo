package compositedatatypes

import (
	"fmt"
	"reflect"
)

// TestSlice - tests slice functionality
func TestSlice() {
	var slice = []int{10, 20, 30} //If you have some starting values, or if a slice’s values aren’t going to change, then a slice literal is a good choice
	var sliceTwo []int            // use this declaration if there is chance that slice won't grow at all.
	//var x = []int{} is equivalent to make([]int, 0) ; non-nil empty slice . The only situation where a zero-length slice is useful is when converting a slice to JSON

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

	// make is used if you know the capacity but not the values
	//pitfall to avoid
	sliceFour := make([]int, 5) //if using slice as a buffer then specify a non-zero length.
	sliceFour = append(sliceFour, 10)
	fmt.Println(sliceFour, len(sliceFour), cap(sliceFour))

	sliceFive := make([]int, 5, 10) //If you are sure you know the exact size you want, you can specify the length and index into the slice to set the values. This is often done when
	//transforming values in one slice and storing them in a second. The downside to this approach is that if you have the size wrong, you’ll end up with either zero values at the end
	//of the slice or a panic from trying to access elements that don’t exist.
	fmt.Println(sliceFive, len(sliceFive), cap(sliceFive))

	//length is zero but it is not nil since memory has been allocated
	//make([]int, 0) is also not nil
	//make([]int, 0, 0) is a compile time error
	sliceSix := make([]int, 0)
	fmt.Println(sliceSix, len(sliceSix), cap(sliceSix))
	fmt.Println(sliceSix == nil)

	sliceSeven := make([]int, 0, 10) //In other situations, use make with a zero-length length and a specified capacity. This allows you to use append to add items to the slice. If the
	//number of items turns out to be smaller, you won’t have an extraneous zero value at the end. If the number of items is larger, your code will not panic.
	sliceSeven = append(sliceSeven, 5, 6, 7, 8)
	fmt.Println(sliceSeven, len(sliceSeven), cap(sliceSeven))

	//x := make([]int, 2, 1) - compile time error to assign capacity less than length

	//slicing
	sliceThree := sparseSlice[1:6] // index:len // this line creates a new slice
	fmt.Println(sliceThree)
	//sliceThree = sparseSlice[:]  // take the whole slice
	//sliceThree = sparseSlice[:3] // start from 0 index till the length provided
	//sliceThree = sparseSlice[2:] // from the 2nd index till the end
	fmt.Println(sliceThree)
}

// SlicingExample - one more example of memory sharing by slicing
func SlicingExample() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	InspectStringSlice("x", x)
	InspectStringSlice("y", y)
	InspectStringSlice("z", z)
	InspectStringSlice("d", d)
	InspectStringSlice("e", e)
}

// Inspect - inspects underlying array for length and capacity and memory address
func Inspect(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}

// SliceMemorySharing -  Slices share storage
func SliceMemorySharing() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10
	z[1] = 30
	Inspect("x", x)
	Inspect("y", y)
	Inspect("z", z)
}

// SliceAppendMemorySharing - impact of append
func SliceAppendMemorySharing() {
	// Whenever you take a slice from another slice, the subslice’s capacity is set to the capacity of the original slice, minus the offset of the subslice within the original slice.
	// This means that any unused capacity in the original slice is also shared with any subslices.
	x := make([]int, 5, 20)
	y := x[3:4]
	Inspect("x", x)
	Inspect("y", y)
	y = append(y, 30)
	Inspect("x", x)
	Inspect("y", y)
}

// ConfusingSlices - For better understanding
func ConfusingSlices() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	Inspect("x", x)
	Inspect("y", y)
	Inspect("z", z)
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	Inspect("x", x)
	Inspect("y", y)
	Inspect("z", z)
}

// FullSliceExpression - for safeguarding capacity limit of subslice, saves the day only from append changes not from updation of array changes
func FullSliceExpression() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	Inspect("x", x)
	Inspect("y", y)
	Inspect("z", z)
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	Inspect("x", x)
	Inspect("y", y)
	Inspect("z", z)
}

// SliceFromArray - slicing of array is similar to slicing from slice (memory sharing)
func SliceFromArray() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	x[3] = 20
	fmt.Println("x", len(x), x, cap(x))
	Inspect("y", y)
	Inspect("z", z)
}

// CopyUsage - copy is the only solution to allocate different memory to slice, saves from append and update changes
func CopyUsage() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	Inspect("x", x)
	Inspect("y", y)
	fmt.Println(num)

	num = copy(y, x[2:])
	Inspect("x", x)
	Inspect("y", y)
	fmt.Println(num)

	num = copy(x[:3], x[1:])
	Inspect("x", x)
	Inspect("y", y)
	fmt.Println(num)

	x = []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y = make([]int, 2)

	copy(y, d[:])
	Inspect("x", x)
	Inspect("y", y)
	fmt.Println("d", len(d), d, cap(d))

	copy(d[:], x)
	Inspect("x", x)
	Inspect("y", y)
	fmt.Println("d", len(d), d, cap(d))
}

// InspectStringSlice - inspects underlying array for length and capacity and memory address
func InspectStringSlice(name string, slice []string) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}
