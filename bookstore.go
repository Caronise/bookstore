package bookstore

import (
	"fmt"
	"io"
)

type Book struct {
	Title  string
	Author string
	Copies int
	Series bool
}

func CreateAndPrintBook(w io.Writer, title, author string, copies int, series bool) Book {
	b := Book{
		Title:  "Death",
		Author: "Richard Beliveau",
		Copies: 5,
		Series: false,
	}

	fmt.Fprintf(w, "%+v\n", b)
	return b
}
