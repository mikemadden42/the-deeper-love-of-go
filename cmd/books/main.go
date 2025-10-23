package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func printBook(book Book) {
	fmt.Printf("%v by %v - %v copies\n", book.Title, book.Author, book.Copies)
}
func main() {
	book := Book{
		Title:  "The Deeper Love of Go",
		Author: "John Arundel",
		Copies: 15,
	}

	fmt.Println("Books in stock:")
	printBook(book)
}
