package main

import (
	"reflect"
	"testing"
)

func TestAddBook(t *testing.T) {
	// ARRANGE
	var store bookStore
	expected := book{
		author: "Marshall B. Rosenberg",
		title:  "nonviolent communication",
		score:  5,
		price:  15,
	}

	// ACT
	store.addBook(expected)

	// ASSERT
	actual := store.books[0]

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual book does not match expected book. Expected: %+v, Actual: %+v", expected, actual)
	}
}

func TestGetBooksWithScoreHigherThan3(t *testing.T) {
	// ARRANGE
	var store bookStore
	bookWithScoreLowerThan3 := book{
		author: "Stephen King",
		title:  "the long walk",
		score:  2,
		price:  8,
	}
	bookWithScoreHigherThan3 := book{
		author: "Tricia Hersey",
		title:  "rest is resistance",
		score:  5,
		price:  12,
	}
	expectedBook := bookWithScoreHigherThan3
	expectedScore := bookWithScoreHigherThan3.score

	// ACT
	store.addBook(bookWithScoreLowerThan3)
	store.addBook(bookWithScoreHigherThan3)
	booksWithScoreHigherThan3 := store.getBooksWithScoreHigherThan(3)

	// ASSERT
	actualBook := booksWithScoreHigherThan3[0]
	actualScore := actualBook.score

	if !reflect.DeepEqual(actualBook, expectedBook) {
		t.Errorf("Actual book does not match expected book. Expected: %+v, Actual: %+v", expectedBook, actualBook)
	}
	if expectedScore != actualScore {
		t.Errorf("Actual book score does not match. Expected: %+v, Actual: %+v", expectedScore, actualScore)
	}
}

func TestTotalPrice(t *testing.T) {
	// ARRANGE
	var store bookStore
	book1 := book{
		author: "Stephen King",
		title:  "the long walk",
		score:  2,
		price:  8,
	}
	book2 := book{
		author: "Tricia Hersey",
		title:  "rest is resistance",
		score:  5,
		price:  12,
	}
	book3 := book{
		author: "Marshall B. Rosenberg",
		title:  "nonviolent communication",
		score:  5,
		price:  15,
	}
	expectedTotalPrice := book1.price + book2.price + book3.price

	store.addBook(book1)
	store.addBook(book2)
	store.addBook(book3)

	// ACT
	actualTotalPrice := store.getTotalPrice()

	// ASSERT
	if expectedTotalPrice != actualTotalPrice {
		t.Errorf("Actual total price does not match. Expected: %+v, Actual: %+v", expectedTotalPrice, actualTotalPrice)
	}
}
