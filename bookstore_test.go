package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getSharedCatalogVar() bookstore.Catalog {
	return bookstore.Catalog{
		{
			ID:          "Book01",
			Title:       "Spark Joy",
			Authors:     []string{"Marie Kondo"},
			Description: "A tiny, cheerful Japanese woman explains tidying.",
			// category:        "Fantasy",
			Copies:     66,
			PriceCents: 1199,
			Series:     false,
		},
		{
			ID:          "Book02",
			Title:       "Death",
			Authors:     []string{"Richard Beliveau"},
			Description: "A deep dive into the mysteries of life and the inevitable reality of death.",
			Copies:      9,
			PriceCents:  2999,
			Series:      false,
		},
		{
			ID:          "Book03",
			Title:       "Lord of the Rings",
			Authors:     []string{"J.R.R Tolkien"},
			Description: "An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.",
			Copies:      20,
			PriceCents:  1950,
			Series:      true,
		},
	}
}

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		ID:          "Book01",
		Title:       "Spark Joy",
		Authors:     []string{"Marie Kondo"},
		Description: "A tiny, cheerful Japanese woman explains tidying.",
		Copies:      66,
		PriceCents:  1199,
		Series:      false,
	}
}

func TestAddBook(t *testing.T) {
	t.Parallel()
	c := getSharedCatalogVar()
	newBook := bookstore.Book{ID: "Book04", Title: "The Elements of Style", Authors: []string{"Strunk", "&", "White"}, Copies: 4, PriceCents: 4999, Series: false}
	c.AddBook(newBook)

	if !cmp.Equal(c[len(c)-1], newBook, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("Failed to add book: %+v to catalog: %+v", newBook, c)
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	c := getSharedCatalogVar()
	got := c.GetAllBooks()

	if !cmp.Equal(c, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(c, got))
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()
	c := getSharedCatalogVar()
	want := "Death by Richard Beliveau - A deep dive into the mysteries of life and the inevitable reality of death.\n"
	got := c.GetBookDetails("Book02")

	if got != want {
		t.Errorf("Expected: \n%q, but got: \n%q", want, got)
	}
}

func TestGetAllBookDetails(t *testing.T) {
	t.Parallel()
	c := getSharedCatalogVar()
	want := "Spark Joy by Marie Kondo - A tiny, cheerful Japanese woman explains tidying.\n" +
		"Death by Richard Beliveau - A deep dive into the mysteries of life and the inevitable reality of death.\n" +
		"Lord of the Rings by J.R.R Tolkien - An epic fantasy novel that chronicles the adventures of hobbits, elves, and men against the dark lord Sauron.\n"
	got := c.GetAllBookDetails()

	if want != got {
		t.Errorf("Expected: \n%q, but got: \n%q", want, got)
	}
}

func TestBuyBook(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	want := "Book01"
	book := c.GetBookByID(want)
	// Change this to fail the purchase
	// book.PriceCents = 0

	purchased := bookstore.BuyBook(book)

	if purchased {
		t.Errorf("Could not purchase %s, copies available: %d", book.Title, book.Copies)
	}
	// fmt.Printf("Succesfully purchased %s for %d\n", book.Title, book.PriceCents)
	// book.Copies -= 1
	// bookstore.Books["Book1"] = book
}

func TestSalePriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "A Clockwork Orange Soda",
		PriceCents: 500,
	}

	want := 500
	got := b.SalePriceCents()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMailingLabel(t *testing.T) {
	t.Parallel()

	c := bookstore.Customer{
		Title:   "Dr.",
		Name:    "Doom",
		Address: "666 End of the Road",
	}

	want := "666 End of the Road"
	got := c.MailingLabel()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestCatalogGetAllBooks(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	got := c.GetAllBooks()

	if !cmp.Equal(c, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf(cmp.Diff(c, got))
	}
}

func TestGetCatalogSize(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	want := 3
	got := c.GetCatalogSize()

	if want != got {
		t.Errorf("Wanted: %d, got %d", want, got)
	}
}

func TestGetBookByID(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()

	want := c[0]
	got := c.GetBookByID(c[0].ID) // "Book01"

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("Wanted %+v, but got %+v", want, got)
	}
}

func TestGetAllTitles(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()

	want := []string{"Spark Joy", "Death", "Lord of the Rings"}
	got := c.GetAllTitles()

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("Wanted %q, but got %q", want, got)
	}
}

func TestGetUniqueAuthors(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()

	newBook := bookstore.Book{
		ID:          "Book05",
		Title:       "The Two Towers",
		Authors:     []string{"J.R.R Tolkien"},
		Description: "The second volume of Tolkien's \"The Lord of the Rings\" trilogy. The tale unfolds with growing darkness and challenges.",
		Copies:      13,
		PriceCents:  1950,
		Series:      true,
	}

	c.AddBook(newBook)

	// J.R.R Tolkien should only appear once, despite being twice in the catalog.
	want := []string{"Marie Kondo", "Richard Beliveau", "J.R.R Tolkien"}
	got := c.GetUniqueAuthors()

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("Wanted: %q, but got %q", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	book := c[0]

	want := 999
	book.SetPriceCents(want)

	if want != book.PriceCents {
		t.Errorf("wanted: %d, but got: %d\n", want, book.PriceCents)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	book := c[0]
	category := "Autobiography"
	err := book.SetCategory(category)
	if err != nil {
		t.Errorf("Failed to set category for book: %s", err.Error())
	}
}

func TestSetDiscountPercent(t *testing.T) {
	t.Parallel()

	c := getSharedCatalogVar()
	book := c[0]

	type testCase struct {
		discount, want int
		errExpected    bool
	}

	testCases := []testCase{
		{discount: 50, want: 50, errExpected: false},
		{discount: 100, want: 100, errExpected: false},
		{discount: 0, want: 0, errExpected: false},
		{discount: 101, want: 0, errExpected: true},
		{discount: -1, want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		book.SetDiscountPercent(0)

		err := book.SetDiscountPercent(tc.discount)
		if tc.errExpected && err == nil {
			t.Errorf("expected an error for discount %d but got none", tc.discount)
		} else if !tc.errExpected && err != nil {
			t.Errorf("did not expect an error for discount %d but got: %s", tc.discount, err)
		}

		got := book.DiscountPercent()
		if tc.want != got {
			t.Errorf("for discount: %d, wanted: %d, but got: %d\n", tc.discount, tc.want, got)
		}
	}

}
