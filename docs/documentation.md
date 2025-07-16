# 📚 Console-Based Library Management System (Go)

A simple console-based Library Management System written in **Go** that demonstrates the use of **structs**, **interfaces**, **methods**, **slices**, and **maps**. This project adheres to clean architecture and is organized for scalability and maintainability.

---

## 🧩 Features

- ✅ Add new books to the library  
- ❌ Remove books from the library  
- 📖 Borrow books (if available)  
- 📤 Return borrowed books  
- 🔍 List all available books  
- 👤 List all borrowed books by a specific member  
- 🔁 Persistent members and book tracking via memory maps

---
## 📁 Folder Structure

library_management/
├── main.go # Entry point of the application
├── controllers/
│ └── library_controller.go # Handles console interaction
├── models/
│ ├── book.go # Defines Book struct
│ └── member.go # Defines Member struct
├── services/
│ └── library_service.go # Implements business logic & interface
├── docs/
│ └── documentation.md # Additional documentation (optional)
└── go.mod # Go module definition

## How to Run
```
git clone https://github.com/your-username/library_management.git
cd library_management
go run main.go
```

## Program Overview
====== Library Management System ======
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books by Member
7. Exit
=======================================
Enter your choice:

