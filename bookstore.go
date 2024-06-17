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
	var store bookStore
	store.utilsAddBooks()

	goodBooks := store.getBooksWithScoreHigherThan(3)
	printBooks(goodBooks)
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

func printBooks(books []book) {
	for _, b := range books {
		fmt.Println()
		fmt.Print(b.title)
		fmt.Printf(" (%s)", b.author)
		fmt.Println()
		fmt.Println(b.score)
	}
}

func (bs *bookStore) getTotalPrice() int {
	totalPrice := 0

	for _, b := range bs.books {
		totalPrice += b.price
	}

	return totalPrice
}

func (bs *bookStore) getBooksWithScoreHigherThan(minScore int) []book {
	var filteredBooks []book

	for _, b := range bs.books {
		if b.score > minScore {
			filteredBooks = append(filteredBooks, b)
		}
	}

	return filteredBooks
}

func (bs *bookStore) addBook(b book) {
	bs.books = append(bs.books, b)
}

func (bs *bookStore) utilsAddBooks() {
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

	book3 := book{
		author: "Marshall B. Rosenberg",
		title:  "nonviolent communication",
		score:  5,
		price:  15,
	}

	book4 := book{
		author: "bell hooks",
		title:  "all about love",
		score:  4,
		price:  11,
	}

	book5 := book{
		author: "Susan Cain",
		title:  "quiet",
		score:  3,
		price:  10,
	}

	book6 := book{
		author: "Marc Bracket",
		title:  "permission to feel",
		score:  4,
		price:  14,
	}

	book7 := book{
		author: "Tricia Hersey",
		title:  "rest is resistance",
		score:  5,
		price:  12,
	}

	bs.books = append(bs.books, book1, book2, book3, book4, book5, book6, book7)
}
