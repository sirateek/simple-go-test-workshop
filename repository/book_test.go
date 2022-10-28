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

var _ = Describe("Book Repository Test", Label("integration"), func() {
	Context("Caching", func() {
		BeforeEach(func() {
			redisClient = cache.NewRedisClient(cache.Config{
				Host: fmt.Sprintf("%s:%s", Redis.Host, Redis.Port),
			})
			bookRepository = repository.NewBookRepository("/Users/sirateek/Repositories/Kaset/go/test-go-workshop/book.csv", redisClient)
		})

		It("should set cache correctly if cache is not exists in the redis but exists in csv file", func() {
			// Expect cache is not exists
			_, err := redisClient.Get("BOOK_1").Result()
			Expect(err).To(Equal(redis.Nil))
			if 

			// Call Repository
			result, err := bookRepository.GetBook(1)
			Expect(err).To(BeNil())

			// Check cache
			resultString, _ := redisClient.Get("BOOK_1").Result()

			// Convert string to struct
			var resultStruct model.BookModel
			json.Unmarshal([]byte(resultString), &resultStruct)

			// Expect the return data to be equal to
			Expect(result).To(Equal(resultStruct))
		})
	})
})
