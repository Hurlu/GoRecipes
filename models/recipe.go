package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
)

type Recipe struct {
	Id                   primitive.ObjectID `bson:"_id"`
	AuthorID             primitive.ObjectID
	Name                 string
	Instructions         string
	IngredientIDs        []primitive.ObjectID
	IngredientQuantities []int
	Tags                 []string
	Picture              []string

	Ingredients []Ingredient `bson:"-"`
	Author      User         `bson:"-"`
}

func (r Recipe) GetUrl() string  {
	return "recipe/" + r.Id.Hex()
}

func (r Recipe) DisplayPic() string {
	if len(r.Picture) == 0 {
		return "/static/img/placeholder.jpg"
	}
	return r.Picture[rand.Intn(len(r.Picture))]

}

func (r *Recipe) FillIngredients(db *mongo.Database) {
	ids := bson.A{}
	for _, id := range r.IngredientIDs {
		ids = append(ids, id)
	}
	res, err := db.Collection("Ingredients").Find(context.TODO(),
		bson.D{{"_id", bson.D{{"$in", ids}}}})
	if err != nil {
		log.Fatal(err)
	}
	for res.Next(context.TODO()) {
		var elem Ingredient
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		r.Ingredients = append(r.Ingredients, elem)
	}
	_ = res.Close(context.TODO())
}
