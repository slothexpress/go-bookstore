package main

import (
	"testing"
)

func TestAddBook(t *testing.T) {
	var store bookStore
	expected := book{
		author: "Marshall B. Rosenberg",
		title:  "nonviolent communication",
		score:  7,
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
