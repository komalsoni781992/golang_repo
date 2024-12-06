package main

import "fmt"

// Book - book type with title author and pages
type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) Read(pages int) {
	if b.Pages > pages {
		b.Pages -= pages
	} else {
		b.Pages = 0
	}
}

func main() {
	b := Book{Title: "Title", Author: "Author", Pages: 100}
	fmt.Println(b.Author, b.Title, b.Pages)
	Read(&b, 30)
	fmt.Println(b.Author, b.Title, b.Pages)
	b.Read(20)
	fmt.Println(b.Author, b.Title, b.Pages)
	b.Read(140)
	fmt.Println(b.Author, b.Title, b.Pages)
}

// Read - subtrct pages from bookcount
func Read(b *Book, pages int) {
	if b.Pages > pages {
		b.Pages -= pages
	} else {
		b.Pages = 0
	}
}
