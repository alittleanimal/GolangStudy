// package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	// 值传递
	book.auth = "666"
}

func changeBook2(book *Book) {
	// 值传递
	book.auth = "777"
}

func main() {
	// var a myint = 10
	// fmt.Println("a= ", a)
	// fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
