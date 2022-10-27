package cache

import (
	"github.com/go-redis/redis"
)

// Config redis config
type Config struct {
	Host     string
	Password string
	DB       int
}

func NewRedisClient(config Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})
}
