package repository

import (
	"errors"
	"os"
	"test-go-workshop/model"

	"github.com/go-redis/redis"
	"github.com/gocarina/gocsv"
)

type bookRepository struct {
	bookFilePath string
	cache        *redis.Client
}

type BookRepository interface {
	GetBook(id int) (model.BookModel, error)
}

var ErrBookNotFound = errors.New("book not found")

func NewBookRepository(bookFilePath string, cache *redis.Client) BookRepository {
	return &bookRepository{
		bookFilePath: bookFilePath,
		cache:        cache,
	}
}

func (b *bookRepository) GetBook(id int) (bookModel model.BookModel, err error) {
	// Get Book From Cache.

	// Cache Miss, Search from file.
	in, err := os.Open(b.bookFilePath)
	if err != nil {
		return bookModel, err
	}
	defer in.Close()

	bookData := []*model.BookModel{}
	if err := gocsv.UnmarshalFile(in, &bookData); err != nil {
		return bookModel, nil
	}

	for _, book := range bookData {
		if book.Id == id {
			return *book, nil
		}
	}

	return bookModel, ErrBookNotFound
}
