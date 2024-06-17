/************************
Coding Assignment:
*************************/
//1. Add a new property called "price" in the book struct

//2. Add more books to the bookStore

//3. Print all scores higher than 3

//4. Print the total price for all the books in the store

//5. Bonus: Come up with an extra task of your own choice :)

/*Installation
***********************/
//Install Go: https://go.dev/doc/install
//Getting started: https://go.dev/doc/tutorial/getting-started
//Hint: to run the code: go run bookstore_assignment.go

/*Good Go syntax reference
************************/
//https://gobyexample.com/

package main

import "fmt"

func main() {
	book1 := book{
		author: "Stephen King",
		title:  "the long walk",
		score:  2,
		price:  8,
	}

	book2 := book{
		author: "Andy Weir",
		title:  "the martian",
		score:  5,
		price:  9,
	}

	var store bookStore
	store.addBook(book1)
	store.addBook(book2)

	for _, book := range store.books {
		fmt.Println(book.title)
	}
}

type book struct {
	author string
	title  string
	score  int
	price  int
}

func (b book) Title() string {
	return b.title
}

type bookStore struct {
	books []book
}

func (bs *bookStore) addBook(b book) {
	bs.books = append(bs.books, b)
}
