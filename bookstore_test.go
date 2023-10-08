package bookstore_test

import (
	"bookstore"
	"strings"
	"testing"
)

func TestPrintBook(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	book := bookstore.Book{
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Copies:          5,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	}

	bookstore.PrintBook(&sb, book)

	expectedOutput := "{Title:Death Authors:[Richard Beliveau] Copies:5 PriceCents:2999 DiscountPercent:10 Series:false}\n"

	if sb.String() != expectedOutput {
		t.Errorf("Expected %q, but got %q", expectedOutput, sb.String())
	}
}

func TestNetPrice(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	book := bookstore.Book{
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Copies:          5,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	}

	bookstore.PrintBook(&sb, book)
	// With the values above it should be 2699
	discountedPrice := bookstore.NetPrice(book.PriceCents, book.DiscountPercent)
	// update priceCents and DiscountPercent to reflect applied discount.
	book.PriceCents = discountedPrice
	book.DiscountPercent = 0

	expectedPrice := 2699
	if discountedPrice != expectedPrice {
		t.Errorf("Expected price: %d, but got %d", expectedPrice, discountedPrice)
	}

	bookstore.PrintBook(&sb, book)
}

func TestPrintMultipleBooks(t *testing.T) {
	t.Parallel()
	var sb strings.Builder

	books := []bookstore.Book{
		{Title: "Death", Authors: []string{"Richard Beliveau"}, Copies: 9, PriceCents: 2999, DiscountPercent: 10, Series: false},
		{Title: "Lord of the rings", Authors: []string{"J.R.R Tolkien"}, Copies: 20, PriceCents: 1950, DiscountPercent: 5, Series: true},
		{Title: "100 years of solitude", Authors: []string{"Gabriel Garcia Marquez"}, Copies: 3, PriceCents: 3500, DiscountPercent: 0, Series: false},
		{Title: "The Elements of Style", Authors: []string{"Strunk", "&", "White"}, Copies: 4, PriceCents: 4999, DiscountPercent: 5, Series: false},
	}

	for _, book := range books {
		bookstore.PrintBook(&sb, book)
	}

	expectedOutput := "{Title:Death Authors:[Richard Beliveau] Copies:9 PriceCents:2999 DiscountPercent:10 Series:false}\n" +
		"{Title:Lord of the rings Authors:[J.R.R Tolkien] Copies:20 PriceCents:1950 DiscountPercent:5 Series:true}\n" +
		"{Title:100 years of solitude Authors:[Gabriel Garcia Marquez] Copies:3 PriceCents:3500 DiscountPercent:0 Series:false}\n" +
		"{Title:The Elements of Style Authors:[Strunk & White] Copies:4 PriceCents:4999 DiscountPercent:5 Series:false}\n"

	if sb.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got %q", expectedOutput, sb.String())
	}
}
