package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseUser    = "user"
	collectionUsers = "users"
)

type userMongoDBRepository struct {
	db *mongo.Client
}

func NewUserMongoDBRepository(db *mongo.Client) UserMongoDBRepositoryStore {
	return &userMongoDBRepository{
		db: db,
	}
}
