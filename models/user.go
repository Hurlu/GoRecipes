package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID`bson:"_id"`
	Username string
	PassHash string
	Salt string
}