package dbops

import (
	"GoRecipes/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect(uri string) *mongo.Database{
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("GoRecipes")
}

func DbIsEmpty(db *mongo.Database) bool {
	nb, _ := db.Collection("Units").CountDocuments(context.TODO(), bson.D{})
	return nb == 0
}

func InitWithBaseData(db *mongo.Database, fixtures bool) {
	if fixtures {
		FixturesFill(db)
	} else {
		defaultUnits := []interface{}{
			models.Unit{Name: "Grams", Units: []string{"mg", "g", "kg"},
				Quantities:[]int{1, 1000, 1000000}},
			models.Unit{Name: "Liters", Units: []string{"mL", "cL", "L"},
				Quantities:[]int{1, 10, 1000}},
		}
		_ , err := db.Collection("Units").InsertMany(context.TODO(), defaultUnits)

		if err != nil{
			log.Fatal(err)
		}
	}
}