package controller

import (
	"address_bloc/model"
	"fmt"
	"os"
	"strconv"
)

type Menu struct {
	AddressBook *model.AddressBook
}

func (menuController *Menu) MainMenu() {
	totalEntries := strconv.Itoa(len(menuController.AddressBook.Entries))

	fmt.Println("Main Menu - " + totalEntries + " entries")
	fmt.Println("1 - View all entries")
	fmt.Println("2 - Create an entry")
	fmt.Println("3 - Search for an entry")
	fmt.Println("4 - Exit")

	var input string
	fmt.Println("Enter your selection: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		menuController.ViewEntries()
	case "2":
		menuController.CreateEntry()
	case "3":
		menuController.SearchEntries()
	case "4":
		fmt.Println("Okie dokie, see ya!")
		os.Exit(0)
	}
}

func (menuController *Menu) ViewEntries() {
	fmt.Println("\n=== View All Entries ===\n")

	for _, entry := range menuController.AddressBook.Entries {
		fmt.Println(entry.Display())
		menuController.EntrySubMenu(entry)
	}

	menuController.MainMenu()
}

func (menuController *Menu) CreateEntry() {
	fmt.Println("\n=== Create New Entry ===\n")

	fmt.Println("Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("\nPhone number: ")
	var phoneNumber string
	fmt.Scanln(&phoneNumber)

	fmt.Println("\nEmail: ")
	var email string
	fmt.Scanln(&email)

	newEntry := &model.Entry{Name: name, PhoneNumber: phoneNumber, Email: email}
	menuController.AddressBook.AddEntry(newEntry)

	fmt.Println("\nNew entry created\n")

	menuController.MainMenu()
}

func (menuController *Menu) EntrySubMenu(entry *model.Entry) {
	fmt.Println("n - next entry")
	fmt.Println("d - delete entry")
	fmt.Println("m - return to main menu")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "n":
		return
	case "d":
		menuController.AddressBook.RemoveEntry(entry.Name, entry.Email, entry.PhoneNumber)
		fmt.Println("\nEntry deleted\n")
		menuController.MainMenu()
	case "m":
		menuController.MainMenu()
	}
}

func (menuController *Menu) SearchEntries() {
	fmt.Println("\n=== Search Entries ===\n")

	var query string
	fmt.Println("Name: ")
	fmt.Scanln(&query)

	entry, err := menuController.AddressBook.Search(query)

	if err != nil {
		fmt.Println("\nEntry not found\n")
		menuController.MainMenu()
	} else {
		fmt.Println(entry.Display())
		menuController.EntrySubMenu(entry)
	}
}
