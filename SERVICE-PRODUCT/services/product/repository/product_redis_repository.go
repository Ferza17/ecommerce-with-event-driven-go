package repository

import "github.com/go-redis/redis/v8"

type productCacheRepository struct {
	client *redis.Client
}

func NewProductCacheRepository(client *redis.Client) ProductCacheRepositoryStore {
	return &productCacheRepository{
		client: client,
	}
}
