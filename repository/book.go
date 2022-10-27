package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"test-go-workshop/cmd"
	"test-go-workshop/model"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
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
	cacheResult, err := b.cache.Get(fmt.Sprint("BOOK_", id)).Result()
	if err == nil {
		logrus.Info("Cache Hit!")
		// Render cache and return
		err = json.Unmarshal([]byte(cacheResult), &bookModel)
		return bookModel, err
	} else {
		logrus.Warn("Warning: Redis Error: ", err)
	}

	// Cache Miss, Search from file.
	in, err := os.Open(b.bookFilePath)
	if err != nil {
		return bookModel, err
	}
	defer in.Close()

	bookStringData, _ := io.ReadAll(in)

	// Read CSV by using our package.
	bookData, err := cmd.ReadCSVFromString(string(bookStringData))
	if err != nil {
		// If error, return it to caller.
		return bookModel, err
	}

	for _, book := range bookData {
		jsonMarshal, err := json.Marshal(book)
		if err == nil {
			// Set JSON cache
			_, err = b.cache.Set(fmt.Sprint("BOOK_", book.Id), string(jsonMarshal), 1*time.Hour).Result()
			if err != nil {
				logrus.Warn("Warning Redis: Cache set Error: ", err)
			}
		}

		if book.Id == id {
			return *book, nil
		}
	}

	return bookModel, ErrBookNotFound
}
