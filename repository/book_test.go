package repository_test

import (
	"fmt"
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

			// Call Repository

			// Check cache

			// Convert string to struct

			// Expect the return data to be equal to
		})
	})
})
