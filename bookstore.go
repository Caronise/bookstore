package bookstore

import (
	"fmt"
	"strings"
)

type Book struct {
	ID              string
	Title           string
	Authors         []string
	Description     string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Series          bool
}

var Books = []Book{
	{
		ID:              "Book01",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondo"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          66,
		PriceCents:      1199,
		DiscountPercent: 10,
		Series:          false,
	},
	{
		ID:              "Book02",
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
		Copies:          9,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	},
	{
		ID:              "Book03",
		Title:           "Lord of the rings",
		Authors:         []string{"J.R.R Tolkien"},
		Description:     "An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.",
		Copies:          20,
		PriceCents:      1950,
		DiscountPercent: 5,
		Series:          true,
	},
}

func GetBookDetails(ID string) string {
	for _, book := range Books {
		if book.ID == ID {
			authors := strings.Join(book.Authors, " ")
			return fmt.Sprintf("%s by %s - %s\n", book.Title, authors, book.Description)
		}
	}
	return ""
}

func ApplyDiscount(priceCents, discountPercent int) int {
	return (priceCents * (100 - discountPercent)) / 100
}

func AddBook(books []Book, newBook Book) []Book {
	books = append(books, newBook)
	return books
}

func GetAllBooks() []Book {
	return Books
}

func GetAllBookDetails() string {
	books := GetAllBooks()
	var sb strings.Builder

	for _, book := range books {
		fmt.Fprint(&sb, GetBookDetails(book.ID))
	}
	return sb.String()
}
