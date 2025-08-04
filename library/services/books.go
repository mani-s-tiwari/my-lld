package services

import (
	"sync"
	"sync/atomic"
)

// Book struct represents a book
type Book struct {
	ID     int64
	Title  string
	Author string
}

// BookService manages books in memory
type BookService struct {
	mu    sync.RWMutex
	books map[int64]*Book
}

var currentBookID int64 = 0

func getNextBookID() int64 {
	return atomic.AddInt64(&currentBookID, 1)
}

// NewBookService initializes the service
func NewBookService() *BookService {
	return &BookService{
		books: make(map[int64]*Book),
	}
}

// AddBook adds a new book with a generated ID
func (bs *BookService) AddBook(title, author string) *Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	book := &Book{
		ID:     getNextBookID(),
		Title:  title,
		Author: author,
	}
	bs.books[book.ID] = book
	return book
}

// ListBooks returns all books
func (bs *BookService) ListBooks() []*Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	var list []*Book
	for _, book := range bs.books {
		list = append(list, book)
	}
	return list
}

// GetBook returns a book by ID
func (bs *BookService) GetBook(id int64) (*Book, bool) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	book, exists := bs.books[id]
	return book, exists
}
