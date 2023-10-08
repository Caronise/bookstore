package main

import (
	"bookstore"
	"fmt"
)

func main() {
	c := bookstore.Catalog{
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

	fmt.Println(c.GetAllBooks())
}
