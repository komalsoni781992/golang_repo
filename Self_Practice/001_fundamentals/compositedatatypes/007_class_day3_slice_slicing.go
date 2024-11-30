package compositedatatypes

import "fmt"

// SliceSlicing - slicing introduction https://github.com/Diwakarsingh052/genspark/blob/main/4-arrays-slices/4-slicing.go
func SliceSlicing() {
	a := []int{10, 20, 30, 40, 50}

	b := a[1:4] // index:len // this line creates a new slice
	fmt.Println(b)
	//b = a[:]  // take the whole slice
	//b = a[:3] // start from 0 index till the length provided
	b = a[2:] // from the 2nd index till the end
	fmt.Println(b)
}
