package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func printBook(book Book) {
	fmt.Println(book.Title, "by", book.Author, "-", book.Copies, "copies")
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
