package controlstructures

import "fmt"

// TestSwitch - https://github.com/Diwakarsingh052/genspark/blob/main/2-control-flow/switch.go
func TestSwitch() {
	ext := ".txt"
	// break is implicit
	// write fallthrough when needed a behaviour of going to the next case
	switch ext {
	case ".txt":
		fmt.Println("a text file")
		fallthrough
	case ".jpeg":
		fmt.Println("an image")
	default:
		fmt.Println("unknown file type")

	}
}
