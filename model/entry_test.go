package model

import "testing"

func TestDisplay(t *testing.T) {
	entry := Entry{Name: "Test", PhoneNumber: "1112223333", Email: "test@example.com"}
	displayableText := entry.Display()

	expectedDisplayableText := "Name: Test\nPhone Number: 1112223333\nEmail: test@example.com\n"

	if displayableText != expectedDisplayableText {
		t.Error("Expected "+expectedDisplayableText+" , got ", displayableText)
	}
}
