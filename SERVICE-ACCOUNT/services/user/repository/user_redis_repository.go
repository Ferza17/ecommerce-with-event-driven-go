package repository

import "github.com/go-redis/redis/v8"

type userCacheRepository struct {
	client *redis.Client
}

func NewUserCacheRepository(client *redis.Client) UserCacheRepositoryStore {
	return &userCacheRepository{
		client: client,
	}
}
