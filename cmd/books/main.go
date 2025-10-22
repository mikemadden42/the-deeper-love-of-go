package main

import "fmt"

func printBook(title, author string) {
	fmt.Println(title, "by", author)
}
func main() {
	fmt.Println("Books in stock:")
	title := "The Deeper Love of Go"
	author := "John Arundel"
	printBook(title, author)
}
