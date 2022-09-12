package bson

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	UserId      string              `bson:"user_id"`
	TotalPrice  float64             `bson:"total_price"`
	CartItems   []CartItem          `bson:"cart_items"`
	CreatedAt   primitive.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at,omitempty"`
	DiscardedAt primitive.Timestamp `bson:"discarded_at,omitempty"`
}

type CartItem struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	ProductId   string              `bson:"product_id"`
	Quantity    int64               `bson:"quantity"`
	Price       float64             `bson:"price"`
	Note        string              `bson:"note"`
	CreatedAt   primitive.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at,omitempty"`
	DiscardedAt primitive.Timestamp `bson:"discarded_at,omitempty"`
}
