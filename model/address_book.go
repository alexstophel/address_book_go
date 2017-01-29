package model

type AddressBook struct {
	Entries []*Entry
}

func (addressBook *AddressBook) AddEntry(entry *Entry) {
	addressBook.Entries = append(addressBook.Entries, entry)
}

func (addressBook *AddressBook) RemoveEntry(name string, email string, phoneNumber string) {
	var indexToRemove int

	entries := addressBook.Entries

	for index, entry := range entries {
		if entry.Name == name && entry.PhoneNumber == phoneNumber && entry.Email == email {
			indexToRemove = index
		}
	}

	addressBook.Entries = append(entries[:indexToRemove], entries[indexToRemove+1:]...)
}
