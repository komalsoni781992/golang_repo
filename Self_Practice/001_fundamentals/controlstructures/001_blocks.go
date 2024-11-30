package controlstructures

import (
	"fmt"
	"log"
)

//Variables, constants, types and functions declared outside of any functions are placed in the package block
//names for other packages that are valid for the file that contains the import statement are in file block

// Shadow - tests how shadow works
func Shadow() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		//x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}

// ShadowWithMultipleAssignment - it is possible with multiple assignment also
func ShadowWithMultipleAssignment() {
	x := 10
	if x > 5 {
		//x, y := 5, 20
		y := 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
}

// ShadowPackageImport - package imported can be shadowed
func ShadowPackageImport() {
	x := 10
	fmt.Println(x)
	fmt := "oops"
	log.Println(fmt)
	//fmt.Println(fmt) - compile time error
	//shadow tool does not detect package import shadow
}

//25 keywords - int, string, true, false, make, close, nil have universe scope
//Go considers these predeclared identifiers and defines them in the universe block, which is the block that contains all other blocks.

// ShadowUniverseBlock -keywords and identifiers in universe scope can be shadowed without shadow tool detecting it.
func ShadowUniverseBlock() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)
}
