package model

import "testing"

func TestAddEntry(t *testing.T) {
	addressBook := AddressBook{}
	entry := &Entry{Name: "Test", PhoneNumber: "1112223333", Email: "test@example.com"}

	addressBook.AddEntry(entry)

	if len(addressBook.Entries) != 1 {
		t.Error("Expected 1, got ", len(addressBook.Entries))
	}
}

func TestRemoveEntry(t *testing.T) {
	addressBook := AddressBook{}
	entry1 := &Entry{Name: "Test1", PhoneNumber: "1112223333", Email: "test1@example.com"}
	entry2 := &Entry{Name: "Test2", PhoneNumber: "1112223333", Email: "test2@example.com"}

	addressBook.AddEntry(entry1)
	addressBook.AddEntry(entry2)

	addressBook.RemoveEntry(entry2.Name, entry2.PhoneNumber, entry2.Email)

	if len(addressBook.Entries) != 1 {
		t.Error("Expected 1, got ", len(addressBook.Entries))
	}
}
