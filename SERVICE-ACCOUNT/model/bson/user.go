package bson

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	Username    string              `bson:"username,omitempty"`
	Email       string              `bson:"email,omitempty"`
	Password    string              `bson:"password,omitempty"`
	Devices     []Device            `bson:"devices,omitempty"`
	CreatedAt   primitive.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at,omitempty"`
	DiscardedAt primitive.Timestamp `bson:"discarded_at,omitempty"`
}

type Device struct {
	Id          primitive.ObjectID  `bson:"id,omitempty"`
	DeviceID    string              `bson:"device_id,omitempty"`
	AccessToken string              `bson:"access_token,omitempty"`
	CreatedAt   primitive.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at,omitempty"`
	DiscardedAt primitive.Timestamp `bson:"discarded_at,omitempty"`
}
