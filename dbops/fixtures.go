package dbops

import (
	"GoRecipes/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var fUsers []models.User

var fUnits []models.Unit

var fIngredients []models.Ingredient

var fRecipes []models.Recipe

func FixturesFill(db *mongo.Database) {
	fUsers = []models.User{{Username: "Paitailservix", Id: primitive.NewObjectID()},
		{Username: "Crocodile Dundee", Id: primitive.NewObjectID()},
		{Username: "Mickaël Wiart ft. Salut c'est cool", Id: primitive.NewObjectID()},
		{Username: "Création Carrefour", Id: primitive.NewObjectID()}}

	fUnits = []models.Unit{
		{Name: "Grammes", Units: []string{"mg", "g", "kg"},
			Quantities: []int{1, 1000, 1000000}, Id: primitive.NewObjectID(), Format: "$QTY $UNIT de $OBJ"},
		{Name: "Litres", Units: []string{"mL", "cL", "L"},
			Quantities: []int{1, 10, 1000}, Id: primitive.NewObjectID(), Format: "$QTY $UNIT de $OBJ"}}

	fIngredients = []models.Ingredient{
		{Name: "Vitriol", Id: primitive.NewObjectID(), UnitID: fUnits[1].Id},
		{Name: "Pétrole", Id: primitive.NewObjectID(), UnitID: fUnits[1].Id},
		{Name: "Viande de crocodile", Id: primitive.NewObjectID(), UnitID: fUnits[0].Id},
		{Name: "Sel de guérande", Id: primitive.NewObjectID(), UnitID: fUnits[0].Id},
		{Name: "Flocons Mousseline", Id: primitive.NewObjectID(), UnitID: fUnits[0].Id},
		{Name: "Lait", Id: primitive.NewObjectID(), UnitID: fUnits[1].Id},
		{Name: "Pâte à tartes", Id: primitive.NewObjectID(), UnitID: fUnits[0].Id},
		{Name: "Tourbe fraîche", Id: primitive.NewObjectID(), UnitID: fUnits[1].Id}}

	fRecipes = []models.Recipe{
		{Name: "Le pudding à l'arsenic",
			Instructions: "Dans un grand bol de strychine [...]",
			Tags:         []string{"Obélix", "Astérix", "Gaulois", "Pâtisserie"}, Id: primitive.NewObjectID(),
			AuthorID:             fUsers[0].Id,
			IngredientIDs:        []primitive.ObjectID{fIngredients[0].Id, fIngredients[1].Id},
			IngredientQuantities: []int{1, 10}},

		{Name: "Steak de crocodile", Instructions: "Crocodile, coupe-deux, barbecue, grill grill",
			Tags: []string{"Dundee", "BBQ"}, Id: primitive.NewObjectID(),
			AuthorID:             fUsers[1].Id,
			IngredientIDs:        []primitive.ObjectID{fIngredients[2].Id, fIngredients[3].Id},
			IngredientQuantities: []int{1, 1000}},

		{Name: "Une bonne purée", Instructions: "Alors pour faire une bonne purée, il faut: Un DPR...",
			Tags: []string{"Classique", "Basique"}, Id: primitive.NewObjectID(),
			AuthorID:             fUsers[2].Id,
			IngredientIDs:        []primitive.ObjectID{fIngredients[4].Id, fIngredients[5].Id},
			IngredientQuantities: []int{98513, 2}},

		{Name: "Tarte à la tourbe", Instructions: "N'oubliez pas de précuire la tourbe avant d'y mettre la pâte !",
			Tags: []string{"Pâtisserie", "Immonde"}, Id: primitive.NewObjectID(),
			AuthorID:             fUsers[3].Id,
			IngredientIDs:        []primitive.ObjectID{fIngredients[6].Id, fIngredients[7].Id},
			IngredientQuantities: []int{32, 9999999999999999}}}

	unit_arr := make([]interface{}, len(fUnits))
	user_arr := make([]interface{}, len(fUsers))
	ingr_arr := make([]interface{}, len(fIngredients))
	recp_arr := make([]interface{}, len(fRecipes))

	for i := range fUnits {
		unit_arr[i] = fUnits[i]
	}

	for i := range fUsers {
		user_arr[i] = fUsers[i]
	}

	for i := range fIngredients {
		ingr_arr[i] = fIngredients[i]
	}

	for i := range fRecipes {
		recp_arr[i] = fRecipes[i]
	}
	errs := make([]error, 4)
	_, errs[0] = db.Collection("Units").InsertMany(context.TODO(), unit_arr)
	_, errs[1] = db.Collection("Users").InsertMany(context.TODO(), user_arr)
	_, errs[2] = db.Collection("Ingredients").InsertMany(context.TODO(), ingr_arr)
	_, errs[3] = db.Collection("Recipes").InsertMany(context.TODO(), recp_arr)

	for i := range errs {
		if errs[i] != nil {
			log.Fatal(errs[i])
		}
	}

}
