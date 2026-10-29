package models

type User struct {
	ID    int
	Name  string
	Email string
}

var Users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}
