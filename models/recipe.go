package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	Id primitive.ObjectID `bson:"_id"`
	AuthorID primitive.ObjectID
	Name string
	Instructions string
	IngredientIDs []primitive.ObjectID
	IngredientQuantities []int
	Tags []string
	Picture [][]byte
}
