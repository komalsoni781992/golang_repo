package basicdatatypes

import "fmt"

// TestNormalAndRawString - raw string and normal string https://github.com/Diwakarsingh052/genspark/blob/main/1-variables/3-raw-string.go
func TestNormalAndRawString() {
	jsonString := "{\n      \"name\": \"Alice\",\n      \"age\": 25\n    }"
	// backtick creates a raw string, no processing happens inside a raw string
	rawString := `{
      "name": "Alice",
      "age": 25
    }`
	filePath := "C:\\Users\\Alice\\Documents\\example.txt"
	rawFilePath := `C:\Users\Alice\Documents\example.txt`

	fmt.Println(jsonString)
	fmt.Println("File Path:", filePath)

	fmt.Println(rawString)
	fmt.Println("File Path:", rawFilePath)
}
