package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exist := l.Books[bookID]

	if !exist || book.Status == "Borrowed" {
		return errors.New("book not available")
	}
	member, exist := l.Members[memberID]

	if !exist {
		return errors.New("member not found")
	}
	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exist := l.Books[bookID]
	if !exist {
		return errors.New("book not found")
	}
	member, exist := l.Members[memberID]
	if !exist {
		return errors.New("member not found")
	}
	newList := []models.Book{}
	found := false

	for _, b := range member.BorrowedBooks {
		if b.ID != bookID {
			newList = append(newList, b)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("book not borrowed by member")
	}
	book.Status = "Available"
	member.BorrowedBooks = newList
	l.Books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, b := range l.Books {
		if b.Status == "Available" {
			available = append(available, b)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exist := l.Members[memberID]
	if !exist {
		return nil
	}
	return member.BorrowedBooks
}
