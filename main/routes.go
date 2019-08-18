package main

import (
	"GoRecipes/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func mainRoutes(router *gin.Engine, db *mongo.Database) {
	recipes := db.Collection("Recipes")

	router.GET("/home", func(c *gin.Context) {
		var objRecipes []*models.Recipe
		cur, err := recipes.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var elem models.Recipe
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			elem.FillIngredients(db)
			for idx := range elem.Ingredients {
				elem.Ingredients[idx].FillUnit(db)
			}
			objRecipes = append(objRecipes, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// Close the cursor once finished
		_ = cur.Close(context.TODO())

		c.HTML(http.StatusOK, "index.html", gin.H{"recipes": objRecipes})
	})

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	router.GET("/tags/:tag", func (c * gin.Context){
		c.HTML(http.StatusOK, "tags.html", gin.H{})
	})

	router.GET("/recipe/:id", func(c *gin.Context) {
		var recipe models.Recipe
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			return
		}
		err = db.Collection("Recipes").FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&recipe)
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "recipe.html", gin.H{"recipe": recipe})
	})

}
