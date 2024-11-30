package basicdatatypes

import "fmt"

// TestRuneLiteral tests rune character data type for Go
func TestRuneLiteral() {
	// same as unit32 internally (alias)
	runeOne := 'a'           //single unicode characters
	runeTwo := '\141'        // 8 bit octal number
	runeThree := '\x61'      // 8 bit hexadecimal number
	runeFour := '\u0061'     // 16 bit unicode number
	runeFive := '\U00000061' // 32 bit unicode numbers
	fmt.Println(runeOne)
	fmt.Println(runeTwo)
	fmt.Println(runeThree)
	fmt.Println(runeFour)
	fmt.Println(runeFive)
	//fmt.Println("\n") //compile time error
	fmt.Println("\t")
	fmt.Println("\\")
	fmt.Println("\"")
	//fmt.Println("\'") //compile time error
}
