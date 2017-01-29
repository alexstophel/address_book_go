package main

import (
	"address_bloc/model"
	"fmt"
)

func main() {
	addressBook := model.AddressBook{}
	testEntry1 := &model.Entry{Name: "Alex", Email: "alex@example.com", PhoneNumber: "1112223333"}
	testEntry2 := &model.Entry{Name: "Jessica", Email: "jessica@example.com", PhoneNumber: "1112223333"}

	addressBook.AddEntry(testEntry1)
	addressBook.AddEntry(testEntry2)

	for _, entry := range addressBook.Entries {
		fmt.Println(entry.Display())
	}
}
