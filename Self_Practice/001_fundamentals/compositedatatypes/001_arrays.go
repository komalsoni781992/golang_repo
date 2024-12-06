package compositedatatypes

import "fmt"

// TestArrayInitialization - tests array in Go
func TestArrayInitialization() {
	var array [3]int
	var arrayInitialized = [3]int{10, 20, 30}
	var sparseArray = [12]int{1, 5: 4, 6, 10: 100, 15} //[1, 0, 0, 0,0, 4, 6, 0, 0, 0, 100, 15].
	var arrayWithoutLength = [...]int{10, 20, 30}
	fmt.Println(array)
	fmt.Println(arrayInitialized)
	fmt.Println(sparseArray)
	fmt.Println(arrayWithoutLength)
	fmt.Println(arrayInitialized == arrayWithoutLength) // type and length should be same
	CompareArrays([3]int{10, 20, 30}, array)
	//the size of the array to be part of the type of the array
	// can’t use a type conversion to convert arrays of different sizes to identical types
	//can’t write a function that works with arrays of any size
	//you can’t assign arrays of different sizes to the same variable.
	//don’t use arrays unless you know the exact length you need ahead of time. For example, some of the cryptographic functions in the
	// standard library return arrays, because the sizes of checksums are defined as part of the algorithm. This is the exception, not the rule.
	for index, value := range sparseArray {
		fmt.Println(index, "value", value)
	}
	fmt.Println(len(sparseArray))
}

// TestMatrix - tests matrix in Go
func TestMatrix() {
	var matrix [2][3]int
	fmt.Println(matrix)
	// matrix[0] = 20 - Invalid
	//An out-of bounds read or write with a variable index compiles but fails at runtime with a panic
	matrix[0] = [3]int{4, 5, 6}
	matrix[0][0] = 50
	fmt.Println(matrix)
	fmt.Println(len(matrix))
}
