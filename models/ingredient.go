package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Ingredient struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string
	UnitID primitive.ObjectID

	Unit Unit `bson:"-"`
}

func (ingredient *Ingredient) FillUnit(db *mongo.Database) {
	res := db.Collection("Units").FindOne(context.TODO(), bson.M{"_id": ingredient.UnitID})
	var elem Unit
	err := res.Decode(&elem)
	if err != nil {
		log.Fatal(err)
	}
	ingredient.Unit = elem
}
