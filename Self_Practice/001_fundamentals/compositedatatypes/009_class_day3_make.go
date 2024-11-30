package compositedatatypes

import (
	"fmt"
)

// MakeIntroduction - https://github.com/Diwakarsingh052/genspark/blob/main/4-arrays-slices/6-make.go
func MakeIntroduction() {
	// make preallocates a backing array
	// very helpful if we have idea about the number of request , we can create size of array according to that
	// in this case append doesn't have to allocate the memory each time
	i := make([]int, 0, 50) // type, len, cap
	Inspect("i", i)
	i = append(i, 100)
	fmt.Println(i)
}
