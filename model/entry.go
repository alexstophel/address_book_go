package model

type Entry struct {
	Email       string
	Name        string
	PhoneNumber string
}

func (entry Entry) Display() string {
	name := "Name: " + entry.Name + "\n"
	phoneNumber := "Phone Number: " + entry.PhoneNumber + "\n"
	email := "Email: " + entry.Email + "\n"

	return name + phoneNumber + email
}
