package bookstore

import (
	"fmt"
	"io"
)

type Book struct {
	Title           string
	Authors         []string
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
