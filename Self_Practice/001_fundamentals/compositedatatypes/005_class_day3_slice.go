package compositedatatypes

import "fmt"

// SliceIntroduction - slices introduction and default nil initialization https://github.com/Diwakarsingh052/genspark/blob/main/4-arrays-slices/2-slice.go
func SliceIntroduction() {
	// slices can grow, there is no limit
	var i []int // create a slice
	//below line would panic, because 0 index doesnt exist, panic is always a runtime error
	// panic crashes the program
	//i[0] = 100 // = is used to update values, not to add values

	if i == nil {
		fmt.Println("i slice is nil")
		fmt.Printf("%#v\n", i)
		return
	}
	fmt.Println(i)
	fmt.Println("end of the main") // unreacahable code
}
