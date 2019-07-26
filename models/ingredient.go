package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	Id primitive.ObjectID`bson:"_id"`
	Name string
	UnitID primitive.ObjectID
}