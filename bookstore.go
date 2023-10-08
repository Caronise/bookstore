package bookstore

import (
	"fmt"
	"io"
)

type Book struct {
	Title           string
	Authors         []string
	Description     string
	Copies          int
	PriceCents      int
	DiscountPercent int
	Series          bool
}

func PrintBook(w io.Writer, book Book) {
	fmt.Fprintf(w, "%+v\n", book)
}

func NetPrice(priceCents, discountPercent int) int {
	return (priceCents * (100 - discountPercent)) / 100
}

func AddBook(books []Book, newBook Book) []Book {
	books = append(books, newBook)
	for i, b := range books {
		fmt.Println(i, ": ", b.Title)
	}
	fmt.Println("Total books after adding new book is:", len(books))
	return books
}
