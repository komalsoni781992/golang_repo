package main

/*q1.	create a program that manages a collection of books, and the number of books available
    allow users to search for books by title.
	The program should handle errors gracefully if a book is not found or if there are any issues accessing the collection.

	Use a map to store book Name and their counter.
	Functionality:
	Implement
		- AddBook(title string,counter int) error
			-to add a new book to the collection.
	FetchBookCounter(name) (int, error)
		    -to retrieve a book by its name.
	Error Handling:
	Use a struct to handle error
	User errors.As in main to check if struct is present inside the chain or not*/
