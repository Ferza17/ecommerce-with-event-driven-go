package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseCart    = "cart"
	collectionCarts = "carts"
)

type cartMongoDBRepository struct {
	db *mongo.Client
}

func NewCartMongoDBRepository(db *mongo.Client) CartMongoDBRepositoryStore {
	return &cartMongoDBRepository{
		db: db,
	}
}
