package bookstore_test

import (
	"bookstore"
	"strings"
	"testing"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondo"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          66,
		PriceCents:      1199,
		DiscountPercent: 10,
		Series:          false,
	}
}

func TestPrintBook(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	book := bookstore.Book{
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
		Copies:          5,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	}

	bookstore.PrintBook(&sb, book)

	expectedOutput := "{Title:Death Authors:[Richard Beliveau] Description:A deep dive into the mysteries of life and the inevitable reality of death. Copies:5 PriceCents:2999 DiscountPercent:10 Series:false}\n"

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
		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
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
		{
			Title:           "Death",
			Authors:         []string{"Richard Beliveau"},
			Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
			Copies:          9,
			PriceCents:      2999,
			DiscountPercent: 10,
			Series:          false,
		},
		{
			Title:           "Lord of the rings",
			Authors:         []string{"J.R.R Tolkien"},
			Description:     "An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.",
			Copies:          20,
			PriceCents:      1950,
			DiscountPercent: 5,
			Series:          true,
		},
		{
			Title:           "100 years of solitude",
			Authors:         []string{"Gabriel Garcia Marquez"},
			Description:     "A magical realist novel that tells the story of the Buendía family in the fictional town of Macondo.",
			Copies:          3,
			PriceCents:      3500,
			DiscountPercent: 0,
			Series:          false,
		},
		{
			Title:           "The Elements of Style",
			Authors:         []string{"Strunk", "&", "White"},
			Description:     "A prescriptive American English writing style guide that provides timeless advice on clarity, simplicity, and usage.",
			Copies:          4,
			PriceCents:      4999,
			DiscountPercent: 5,
			Series:          false,
		},
	}

	for _, book := range books {
		bookstore.PrintBook(&sb, book)
	}

	expectedOutput := "{Title:Death Authors:[Richard Beliveau] Description:A deep dive into the mysteries of life and the inevitable reality of death. Copies:9 PriceCents:2999 DiscountPercent:10 Series:false}\n" +
		"{Title:Lord of the rings Authors:[J.R.R Tolkien] Description:An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron. Copies:20 PriceCents:1950 DiscountPercent:5 Series:true}\n" +
		"{Title:100 years of solitude Authors:[Gabriel Garcia Marquez] Description:A magical realist novel that tells the story of the Buendía family in the fictional town of Macondo. Copies:3 PriceCents:3500 DiscountPercent:0 Series:false}\n" +
		"{Title:The Elements of Style Authors:[Strunk & White] Description:A prescriptive American English writing style guide that provides timeless advice on clarity, simplicity, and usage. Copies:4 PriceCents:4999 DiscountPercent:5 Series:false}\n"

	if sb.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got %q", expectedOutput, sb.String())
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()

	books := []bookstore.Book{
		{Title: "Death", Authors: []string{"Richard Beliveau"}, Copies: 9, PriceCents: 2999, DiscountPercent: 10, Series: false},
		{Title: "Lord of the rings", Authors: []string{"J.R.R Tolkien"}, Copies: 20, PriceCents: 1950, DiscountPercent: 5, Series: true},
	}

	newBook := bookstore.Book{Title: "The Elements of Style", Authors: []string{"Strunk", "&", "White"}, Copies: 4, PriceCents: 4999, DiscountPercent: 5, Series: false}

	updatedBooks := bookstore.AddBook(books, newBook)

	if len(updatedBooks) != len(books)+1 {
		t.Errorf("Failed to add book: %+v to books: %+v", newBook, books)
	}
}
