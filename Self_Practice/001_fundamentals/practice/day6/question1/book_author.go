package main

import "fmt"

/*q1. Create a struct (Author)
  Two Field:- Name, Books[slice]
  Create two methods, one appends new books to the book slice , other prints the struct

  Create a function that accepts the struct and append values to the book slice

  Create a function that would accept the Books field, not the struct and append some more books*/

// Author - Author type with name and list of books
type Author struct {
	name  string
	books []string
}

func (a *Author) appendBook(book string) {
	a.books = append(a.books, book)
}

func (a *Author) printAuthor() {
	fmt.Printf("\n%#v\n", *a)
}

func main() {
	a := Author{name: "Komal"}
	a.appendBook("FirstBook")
	appendBookFunction(&a, "SecondBook")
	appendBookFunctionByBooks(&(a.books), "Thirdbook")
	a.printAuthor()
}

func appendBookFunction(a *Author, book string) {
	a.books = append(a.books, book)
}

func appendBookFunctionByBooks(books *[]string, book string) {
	*books = append(*books, book)
}
