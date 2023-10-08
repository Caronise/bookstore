package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		ID:              "Book01",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondo"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          66,
		PriceCents:      1199,
		DiscountPercent: 10,
		Series:          false,
	}
}

func TestApplyDiscount(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		ID:              "Book02",
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
		Copies:          5,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	}

	// In this example with the values above it should be 2699
	expectedPrice := (book.PriceCents * (100 - book.DiscountPercent)) / 100
	// Apply the discount
	discountedPrice := bookstore.ApplyDiscount(book.PriceCents, book.DiscountPercent)
	// update priceCents and DiscountPercent to reflect applied discount.
	book.PriceCents = discountedPrice
	book.DiscountPercent = 0

	if discountedPrice != expectedPrice {
		t.Errorf("Expected price: %d, but got %d", expectedPrice, discountedPrice)
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()

	books := bookstore.Books

	newBook := bookstore.Book{ID: "Book04", Title: "The Elements of Style", Authors: []string{"Strunk", "&", "White"}, Copies: 4, PriceCents: 4999, DiscountPercent: 5, Series: false}

	updatedBooks := bookstore.AddBook(books, newBook)

	if _, exists := updatedBooks[newBook.ID]; !exists {
		t.Errorf("Failed to add book: %+v to books: %+v", newBook, books)
	}
}

func TestGetAllBooks(t *testing.T) {
	want := bookstore.Books

	got := bookstore.GetAllBooks()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Death by Richard Beliveau - A deep dive into the mysteries of life and the inevitable reality of death.\n"

	got := bookstore.GetBookDetails("Book02")

	if got != want {
		t.Errorf("Expected: \n%q, but got: \n%q", want, got)
	}
}

func TestGetAllBookDetails(t *testing.T) {
	t.Parallel()

	want := "Spark Joy by Marie Kondo - A tiny, cheerful Japanese woman explains tidying.\n" +
		"Death by Richard Beliveau - A deep dive into the mysteries of life and the inevitable reality of death.\n" +
		"Lord of the rings by J.R.R Tolkien - An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.\n"

	got := bookstore.GetAllBookDetails()

	if want != got {
		t.Errorf("Expected: \n%q, but got: \n%q", want, got)
	}
}
