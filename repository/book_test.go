package repository_test

import (
	"encoding/json"
	"fmt"
	"test-go-workshop/model"
	cache "test-go-workshop/pkg"
	"test-go-workshop/repository"

	"github.com/go-redis/redis"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var bookRepository repository.BookRepository
var redisClient *redis.Client

var _ = Describe("Repository", Label("integration"), func() {
	Context("Caching Feature", func() {
		BeforeEach(func() {
			redisClient = cache.NewRedisClient(cache.Config{
				Host: fmt.Sprintf("%s:%s", Redis.Host, Redis.Port),
			})
			bookRepository = repository.NewBookRepository("/Users/sirateek/Repositories/Kaset/go/test-go-workshop/book.csv", redisClient)
		})

		It("should set cache if the cache is not exists and the book is found in CSV", func() {
			// New Redis, No Cache.
			bookModel, err := bookRepository.GetBook(1)
			Expect(err).To(BeNil())

			// Check if cache is set correctly.
			result, err := redisClient.Get("BOOK_1").Result()
			Expect(err).To(BeNil())

			var resultBookModel model.BookModel
			json.Unmarshal([]byte(result), &resultBookModel)

			Expect(bookModel.Id).To(Equal(resultBookModel.Id))
			Expect(bookModel.Title).To(Equal(resultBookModel.Title))
			Expect(bookModel.Author).To(Equal(resultBookModel.Author))
			Expect(bookModel.ISBN).To(Equal(resultBookModel.ISBN))
		})
	})
})
