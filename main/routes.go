package main

import (
	"GoRecipes/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func mongoConnect() *mongo.Client{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
	return client
}

func main() {
	router := gin.Default()

	var cli *mongo.Client = mongoConnect()
	var db *mongo.Database = cli.Database("GoRecipes")


	testUnit := models.Unit{Name: "Grams", Units: []string{"mg", "g", "kg"},
		Quantities:[]int{1, 1000, 1000000}, DefaultUnitIndex:1}

	units := db.Collection("Units")

	_ , err := units.InsertOne(context.TODO(), testUnit)
	if err != nil{
		log.Fatal(err)
	}

	auth_routes(router)

	router.NoRoute(func(c * gin.Context){
		c.String(http.StatusOK, "SALUTOULMONDE C'EST DAVID LAFARGE \nP\nO\nK\nE\nM\nO\nN\nAUJOURD'HUI POUR OUVRIR DE NOUVEAUX BOOSTERS!")
	})

	router.GET("/units", func(c *gin.Context) {
		var unit models.Unit
		filter := bson.D{{"name", "Grams"}}
		err := units.FindOne(context.TODO(), filter).Decode(&unit)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, "Unit name: %s. Default symbol : %s", unit.Name, unit.Units[unit.DefaultUnitIndex])
	})

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
			//c.FullPath() == "/user/:name/*action" // true
	})

	router.Run()
}