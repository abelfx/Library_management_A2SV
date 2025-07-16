package models

// Defining the member struct
type Member struct {
	ID int
	Name string
	BorrowedBooks []Book  // a slice to hold borrowed books
}