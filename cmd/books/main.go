package main

import "fmt"

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")
}
func main() {
	fmt.Println("Books in stock:")
	title := "The Deeper Love of Go"
	author := "John Arundel"
	copies := 15
	printBook(title, author, copies)
}
