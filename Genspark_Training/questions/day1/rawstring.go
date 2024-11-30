package main

import "fmt"

/*q5. {
    "name": "Alice",
    "age": 25
  }

  C:\Users\Alice\Documents\example.txt

  Store above data in string*/

func main() {
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
