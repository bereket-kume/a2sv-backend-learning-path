package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func StartConsole(library services.LibraryManager, members map[int]*models.Member) {
	var bookIDCounter int
	for {
		fmt.Println("\nLibrary menu")
		fmt.Println("1: Add new book")
		fmt.Println("2: remove existing book")
		fmt.Println("3: Borrow book")
		fmt.Println("4: Return book")
		fmt.Println("5: List available books")
		fmt.Println("6: List borrowed books")
		fmt.Println("0: exit")
		fmt.Println("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:

			var title, author string
			fmt.Println("Title: ")
			fmt.Scan(&title)
			fmt.Println("Author")
			fmt.Scan(&author)
			book := models.Book{ID: bookIDCounter, Title: title, Author: author, Status: "Available"}
			library.AddBook(book)
			bookIDCounter++
			fmt.Println("Book added successfully")
		case 2:
			var id int
			fmt.Println("Book Id to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
			fmt.Println("Book removed")

		case 3:
			var bookId, memberId int
			fmt.Println("Book Id to borrow: ")
			fmt.Scan(&bookId)
			fmt.Println("Member Id: ")
			fmt.Scan(&memberId)
			err := library.BorrowBook(bookId, memberId)

			if err != nil {
				fmt.Print("Error:", err)
			} else {
				fmt.Println("Book borrowed")
			}
		case 4:
			var bookID, memberID int
			fmt.Println("Book id to return: ")
			fmt.Scan(&bookID)
			fmt.Println("member Id: ")
			fmt.Scan(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Book returned")
			}
		case 5:
			books := library.ListAvailableBooks()
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 6:
			var memberID int
			fmt.Println("Member ID")
			fmt.Scan(&memberID)
			books := library.ListBorrowedBooks(memberID)
			for _, book := range books {
				fmt.Printf("ID: %d, Titlle: %s, Author: %s", book.ID, book.Title, book.Author)
			}
		case 0:
			fmt.Println("Exiting the library system. goodbye")
			return
		default:
			fmt.Println("invalid choice")
		}
	}
}
