package services

import (
	"errors"
	"library-management/models"
)

type LibraryManger interface {
	AddBook(book models.Book)
	BorrowBook(bookId int, memberId int) error
	RemoveBook(bookId int)
	ReturnBook(bookId int, memberId int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
}


type Library struct {
	Books map[int]models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}


func (l *Library) AddBook(book models.Book) {
	l.Books[book.Id] = book
}


func (l *Library) RemoveBook(bookId int) {
	delete(l.Books, bookId)
}

func (l *Library) BorrowBook(bookId int, memberId int) error {
	book, exists := l.Books[bookId]
	if !exists {
		return errors.New("book not found")
	} 

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, exists := l.Members[memberId]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookId] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l * Library) ReturnBook(bookId int, memeberId int) error {
	member, exists := l.Members[memeberId]
	if !exists {
		return  errors.New("member not found")
	}

	for idx, book := range member.BorrowedBooks {
		if book.Id == bookId {
			member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]... )

			book = l.Books[bookId]
			book.Status = "Available"
			l.Books[bookId] = book
			return nil
		}
	}

	return errors.New("book not borrowed by this member")
}

func (l *Library) ListAvailableBooks() []models.Book {
	var books []models.Book

	for _, book := range l.Books {
		if book.Status == "Available" {
			books = append(books, book)
		}
	}
	return books
}

func (l *Library) ListBorrowedBooks(memberId int) []models.Book {
	member, exist := l.Members[memberId]
	if exist {
		return  []models.Book{}
	}
	return  member.BorrowedBooks
}