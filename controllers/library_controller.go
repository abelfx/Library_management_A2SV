package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(lib *services.Library) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			id := readInt(scanner)
			fmt.Print("Enter Title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()
			lib.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
			fmt.Println("Book added.")
		case "2":
			fmt.Print("Enter Book ID to remove: ")
			id := readInt(scanner)
			lib.RemoveBook(id)
			fmt.Println("Book removed.")
		case "3":
			fmt.Print("Enter Book ID: ")
			bookID := readInt(scanner)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(scanner)
			err := lib.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case "4":
			fmt.Print("Enter Book ID: ")
			bookID := readInt(scanner)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(scanner)
			err := lib.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case "5":
			books := lib.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
			}
		case "6":
			fmt.Print("Enter Member ID: ")
			memberID := readInt(scanner)
			books := lib.ListBorrowedBooks(memberID)
			fmt.Printf("Borrowed Books by Member %d:\n", memberID)
			for _, book := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
			}
		case "7":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	input := scanner.Text()
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	return num
}
