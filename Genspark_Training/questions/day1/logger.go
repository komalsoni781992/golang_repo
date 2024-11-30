package main

import (
	"log"
	"os"
)

/*q3. Use log.New() and print a log message
    use shortfile flag and stdFlags for flag argument (bit of Google required)
// Hint:- first value to this function could be os.Stdout
// Printing: l.Println("hello this is first log message")*/

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Lshortfile|log.LstdFlags)
	logger.Println("Hello World")
}
