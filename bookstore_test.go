package bookstore_test

import (
	"bookstore"
	"strings"
	"testing"
)

func TestCreateAndPrintBook(t *testing.T) {
	t.Parallel()
	var sb strings.Builder

	_ = bookstore.CreateAndPrintBook(&sb, "Death", "Richard Beliveau", 5, false)

	expectedOutput := "{Title:Death Author:Richard Beliveau Copies:5 Series:false}\n"

	if sb.String() != expectedOutput {
		t.Errorf("Expected %q, but got %q", expectedOutput, sb.String())
	}
}
