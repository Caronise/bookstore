package bookstore

import (
	"fmt"
	"strings"
)

// Book contains the information for each book.
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

// Catalog is a slice of books.
type Catalog []Book

// Customer contains the information for each customer.
type Customer struct {
	Title   string
	Name    string
	Address string
}

// GetBookDetails returns a string with details on a book in the format
// "title by author - description"
func (c Catalog) GetBookDetails(ID string) string {
	for _, book := range c {
		if book.ID == ID {
			authors := strings.Join(book.Authors, " ")
			return fmt.Sprintf("%s by %s - %s\n", book.Title, authors, book.Description)
		}
	}
	return ""
}

// GetBookByID returns the book based on book ID.
func (c Catalog) GetBookByID(ID string) Book {
	for _, b := range c {
		if b.ID == ID {
			return b
		}
	}
	return Book{}
}

// EvaluateDiscount returns the price after applying the discount.
func EvaluateDiscount(priceCents, discountPercent int) int {
	return (priceCents * (100 - discountPercent)) / 100
}

// AddBook appends a new book to catalog
func (c Catalog) AddBook(newBook Book) Catalog {
	c = append(c, newBook)
	return c
}

// GetAllBookDetails returns a string that concatenates GetBookDetails of
// every book in the catalog.
func (c Catalog) GetAllBookDetails() string {
	catalog := c.GetAllBooks()
	var sb strings.Builder

	for _, b := range catalog {
		fmt.Fprint(&sb, c.GetBookDetails(b.ID))
	}

	return sb.String()
}

// BuyBook returns true if book.Copies is > 0
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

// GetAllBooks returns the slice of all books inside the catalog.
func (c Catalog) GetAllBooks() Catalog {
	return []Book(c)
}

// GetCatalogSize returns the size of the catalog.
func (c Catalog) GetCatalogSize() int {
	return len(c)
}

// GetAllTitles returns a slice of all the book titles inside the catalog.
func (c Catalog) GetAllTitles() []string {
	titles := make([]string, 0, len(c))
	for _, b := range c {
		titles = append(titles, b.Title)
	}
	return titles
}

func contains(s []string, str string) bool {
	for _, val := range s {
		if val == str {
			return true
		}
	}
	return false
}

func (c Catalog) GetUniqueAuthors() []string {
	authors := make([]string, 0)
	allAuthors := make([]string, 0, len(c))
	for _, b := range c {
		allAuthors = append(allAuthors, b.Authors...)
	}

	for _, a := range allAuthors {
		if !contains(authors, a) {
			authors = append(authors, a)
		}
	}

	return authors
}
