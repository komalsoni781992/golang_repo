package basicdatatypes

import "fmt"

// strings are not collection/sequence of runes instead they are made up of bytes
//These byte don’t have to be in any particular character encoding, but several Go library functions (and the for-range loop that we discuss in the next chapter) assume that
// a string is composed of a sequence of UTF-8-encoded code points.

// TestString tests string data type for Go
func TestString() {
	stringOne := "Greetings and\n\"Salutations\"" // normal string
	stringTwo := `Greetings and
	"Salutations"` //raw string
	fmt.Println(stringOne)
	fmt.Println(stringTwo)
	fmt.Println(len(stringOne))
	fmt.Println(len(stringTwo))        // raw string treats 4 characters as 1 tab character
	fmt.Println(stringOne + stringTwo) //concatenation
}

// StringSlicing slicing and indexing works well in strings. Since strings are immutable, they don’t have the modification problems that slices of slices do. There is a different problem, though.
func StringSlicing() {
	var s string = "Hello there"
	var b byte = s[6]
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(b)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

// A string is composed of a sequence of bytes, while a code point in UTF8 can be anywhere from 1 to 4 bytes long.
// Our previous example was entirely composed of code points that are one byte long in UTF-8, so everything worked out as expected.
// But when dealing with languages other than English or with emojis, you run into code points that are multiple bytes long in UTF-8.

// SlicingEncodingIssue - when dealing with languages other than English or with emojis, you run into code points that are multiple bytes long in UTF-8.
func SlicingEncodingIssue() {
	var s string = "Hello Ѝउ"
	fmt.Println(len(s))
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s[6:])
	fmt.Println(s[7:])
	fmt.Println(s[8:])
	fmt.Println(s[9:])
	fmt.Println(s[10:])
}

// RuneByteToStringConversion - rune and byte can be converted to a string
func RuneByteToStringConversion() {
	//Most data in Go is read and written as a sequence of bytes, so the most common string type conversions are back and forth with a slice of bytes.
	// Slices of runes are uncommon.
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#v\n", s2)
	//var x int = 65  //no other type other than rune or byte can be converted to string
	//var y = string(x)
}

// StringToRuneByteSlices - conversion of string to rune or byte slices
func StringToRuneByteSlices() {
	var s string = "Hello,  Ѝउ"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}
