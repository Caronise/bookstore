package bookstore

import (
	"fmt"
	"sort"
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

type Customer struct {
	Title   string
	Name    string
	Address string
}

// var Books = []Book{
// 	{
// 		ID:              "Book01",
// 		Title:           "Spark Joy",
// 		Authors:         []string{"Marie Kondo"},
// 		Description:     "A tiny, cheerful Japanese woman explains tidying.",
// 		Copies:          66,
// 		PriceCents:      1199,
// 		DiscountPercent: 10,
// 		Series:          false,
// 	},
// 	{
// 		ID:              "Book02",
// 		Title:           "Death",
// 		Authors:         []string{"Richard Beliveau"},
// 		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
// 		Copies:          9,
// 		PriceCents:      2999,
// 		DiscountPercent: 10,
// 		Series:          false,
// 	},
// 	{
// 		ID:              "Book03",
// 		Title:           "Lord of the rings",
// 		Authors:         []string{"J.R.R Tolkien"},
// 		Description:     "An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.",
// 		Copies:          20,
// 		PriceCents:      1950,
// 		DiscountPercent: 5,
// 		Series:          true,
// 	},
// }

var Books = map[string]Book{
	"Book01": {
		ID:              "Book01",
		Title:           "Spark Joy",
		Authors:         []string{"Marie Kondo"},
		Description:     "A tiny, cheerful Japanese woman explains tidying.",
		Copies:          66,
		PriceCents:      1199,
		DiscountPercent: 10,
		Series:          false,
	},
	"Book02": {
		ID:              "Book02",
		Title:           "Death",
		Authors:         []string{"Richard Beliveau"},
		Description:     "A deep dive into the mysteries of life and the inevitable reality of death.",
		Copies:          9,
		PriceCents:      2999,
		DiscountPercent: 10,
		Series:          false,
	},
	"Book03": {
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

// // AddBook to slice
// func AddBook(books []Book, newBook Book) []Book {
// 	books = append(books, newBook)
// 	return books
// }

// AddBook to map
func AddBook(books map[string]Book, newBook Book) map[string]Book {
	books[newBook.ID] = newBook
	return books
}

// // GetAllBooks from slice
// func GetAllBooks() []Book {
// 	return Books
// }

// GetAllBooks from map
func GetAllBooks() map[string]Book {
	return Books
}

func GetAllBookDetails() string {
	books := GetAllBooks()
	var sb strings.Builder

	keys := make([]string, 0, len(books))
	for bk := range books {
		keys = append(keys, bk)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprint(&sb, GetBookDetails(k))
	}

	return sb.String()
}

func BuyBook(book Book) bool {
	return book.Copies <= 0
}

// SalePriceCents returns the sale price of the book.
func (b Book) SalePriceCents() int {
	return b.PriceCents
}

// MailingLabel returns the address of a customer.
func (c Customer) MailingLabel() string {
	return c.Address
}
