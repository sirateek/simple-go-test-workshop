package main

import (
	"fmt"
	cache "test-go-workshop/pkg"
	"test-go-workshop/repository"
)

func main() {
	cacheConfig := cache.Config{}
	redisClient := cache.NewRedisClient(cacheConfig)

	bookRepository := repository.NewBookRepository("/Users/sirateek/Repositories/Kaset/go/test-go-workshop/book.csv", redisClient)

	result, err := bookRepository.GetBook(1)
	fmt.Println("Book Result: ", result, "|| Err: ", err)
}
