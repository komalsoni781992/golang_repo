package helloworld

import (
	"fmt"
	"os"
)

// HelloWorld tests Hello World for Go
func HelloWorld() {
	fmt.Println("Hello World!")

	y := 2 + 5 +
		6 // newline could be after + but not before +
	fmt.Println(y)

	fmt.Println("Hello, Ѝउ") //International characters allowed

	fmt.Println("Hello, my name is", "Komal")

	os.Exit(0) // takes status as parameter
}
