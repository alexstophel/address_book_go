package main

import (
	"address_bloc/controller"
	"address_bloc/model"
)

func main() {
	addressBook := &model.AddressBook{}
	entry := &model.Entry{Name: "Alex", Email: "alex@example.com", PhoneNumber: "1112223333"}
	addressBook.AddEntry(entry)

	menuController := controller.Menu{AddressBook: addressBook}

	menuController.MainMenu()
}
