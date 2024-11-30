package controlstructures

import "fmt"

// TestLoops - https://github.com/Diwakarsingh052/genspark/blob/main/2-control-flow/loops.go
func TestLoops() {
	for i, j := 0, 10; i <= 10; i++ {
		fmt.Println(j)
	}

	i := 0
	for ; i <= 10; i++ {
	}

	i = 0
	for i <= 10 {
		i++
	}
}
