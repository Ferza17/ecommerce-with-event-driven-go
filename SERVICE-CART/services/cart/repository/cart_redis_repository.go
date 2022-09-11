package repository

import "github.com/go-redis/redis/v8"

type cartRedisRepository struct {
	client *redis.Client
}

func NewCartRedisRepository(client *redis.Client) CartRedisRepositoryStore {
	return &cartRedisRepository{
		client: client,
	}
}
