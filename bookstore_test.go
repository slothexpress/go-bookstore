package main

import (
	"reflect"
	"testing"
)

func TestAddBook(t *testing.T) {
	var store bookStore
	expected := book{
		author: "Marshall B. Rosenberg",
		title:  "nonviolent communication",
		score:  5,
		price:  15,
	}
	numberOfBooks := 1

	store.addBook(expected)
	actual := store.books[0]

	if len(store.books) != numberOfBooks {
		t.Errorf("Expected %d book. Actual: %d", numberOfBooks, len(store.books))
	}
	if store.books[0] != expected {
		t.Errorf("Actual book does not match expected book. Expected: %+v, Actual: %+v", expected, actual)
	}
}

func TestGetBooksWithScoreHigherThan3(t *testing.T) {
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
	numberOfBooks := 1
	expectedBook := bookWithScoreHigherThan3
	expectedScore := bookWithScoreHigherThan3.score

	store.addBook(bookWithScoreLowerThan3)
	store.addBook(bookWithScoreHigherThan3)
	booksWithScoreHigherThan3 := store.getBooksWithScoreHigherThan(3)
	actualBook := booksWithScoreHigherThan3[0]
	actualScore := actualBook.score

	if len(booksWithScoreHigherThan3) != numberOfBooks {
		t.Errorf("Expected %d book. Actual: %d", numberOfBooks, len(booksWithScoreHigherThan3))
	}
	if !reflect.DeepEqual(actualBook, expectedBook) {
		t.Errorf("Actual book does not match expected book. Expected: %+v, Actual: %+v", expectedBook, actualBook)
	}
	if expectedScore != actualScore {
		t.Errorf("Actual book score does not match. Expected: %+v, Actual: %+v", expectedScore, actualScore)
	}
}
