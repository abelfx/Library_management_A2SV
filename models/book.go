package models

// Defining the Book Struct

type Book struct {
	ID int
	Title string
	Author string
	Status string // can be "Available" or "Borrowed"

}