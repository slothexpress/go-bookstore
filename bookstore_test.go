package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expected, actual, "Expected book does not match actual book.\nExpected: %+v\nActual:   %+v", expected, actual)
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
	minScore := 3
	expected := bookWithScoreHigherThan3

	store.addBook(bookWithScoreLowerThan3)
	store.addBook(bookWithScoreHigherThan3)

	// ACT
	booksWithScoreHigherThan3 := store.getBooksWithScoreHigherThan(minScore)

	// ASSERT
	actual := booksWithScoreHigherThan3[0]
	actualScore := actual.score

	assert.Equal(t, expected, actual,
		"Expected book does not match actual book.\nExpected: %+v\nActual:   %+v", expected, actual)
	assert.True(t, actualScore > minScore,
		"Score is not higher than minimum score.\nMinScore: %+v\nActual:   %+v", minScore, actualScore)
}

func TestGetTotalPrice(t *testing.T) {
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
	expected := book1.price + book2.price + book3.price

	store.addBook(book1)
	store.addBook(book2)
	store.addBook(book3)

	// ACT
	actual := store.getTotalPrice()

	// ASSERT
	assert.Equal(t, expected, actual,
		"Actual total price does not match. Expected: %+v, Actual: %+v", expected, actual)
}
