package main

import (
	"GoRecipes/models"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
	"net/http"
)

func mainRoutes(router *gin.Engine, db *mongo.Database) {
	router.GET("/home", func(c *gin.Context) {
		var objRecipes []models.Recipe
		objRecipes = models.GetAllFilledRecipes(db)
		c.HTML(http.StatusOK, "index.html", gin.H{"recipes": objRecipes})
	})

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
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


	router.POST("/query", func(c *gin.Context) {
		type Query struct {
			Query     string
		}
		var json Query
		c.Bind(&json)
		var objRecipes []models.Recipe
		objRecipes = models.SearchRecipes(json.Query, db)
		tmpl, err := template.New("recipes").ParseFiles("_frontend/html/recipes.html")
		if err != nil { log.Fatal(err) }
		var str_tpl bytes.Buffer
		err = tmpl.Execute(&str_tpl, gin.H{"recipes": objRecipes})
		if err != nil { log.Fatal(err) }
		c.JSON(200, gin.H{"value": str_tpl.String()})
	})
}
