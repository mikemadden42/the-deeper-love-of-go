package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v - %v copies", book.Title, book.Author, book.Copies)
}

func main() {
	book := Book{
		Title:  "The Deeper Love of Go",
		Author: "John Arundel",
		Copies: 15,
	}

	fmt.Println("Books in stock:")
	fmt.Print(BookToString(book))
}
