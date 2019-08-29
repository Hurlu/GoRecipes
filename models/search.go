package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"regexp"
)

func AddNameBsonE(name string) bson.E{
	//bson.D{{"name", bson.D{{"$regex", name}, {"$options", "i"}}}}
	if name != "" {
		return bson.E{"name", bson.D{{"$regex", name}, {"$options", "i"}}}
	}
	return bson.E{}
}

func AddTagBsonE (tags []string) bson.E{
	if len(tags) > 0 {
		return bson.E{"tags",bson.D{{"$in", tags}}}}
	return bson.E{}
}

func AddNegTagBsonE(negtags []string) bson.E{
	if len(negtags) > 0 {
		return bson.E{"tags",bson.D{{"$not", bson.D{{"$in", negtags}}}}}
	}
	return bson.E{}

}

func AddTagSearch(tag string, data string, tags *[]string, negtags *[]string){
	switch tag {
	case "T":
		if data[0] == '-' {
			*negtags = append(*negtags, data[1:])
		} else {
			*tags = append(*tags, data)
		}
		break
	case "I":
		break
	}
}

func QueryStringToMongoQuery(query string) bson.D {
	name := ""
	tags := make([]string, 0)
	negtags := make([]string, 0)
	regexp, err := regexp.Compile(`(\((\p{L}):([\p{L} *-]+)\)|[\p{L} *-]+)`)
	if err != nil { panic(err) }
	matches := regexp.FindAllStringSubmatch(query, -1)
	if matches != nil {
		for _, submatches := range matches {
			if len(name) == 0 && len(submatches[1]) > 0 && len(submatches[2]) == 0 && len(submatches[3]) == 0 {
				name = submatches[1]
			} else if len(submatches[2]) > 0 && len(submatches[3]) > 0 {
				AddTagSearch(submatches[2], submatches[3], &tags, &negtags)
			}
		}
	} else {
		print("No 1st level matches !")
	}
	return bson.D{AddNameBsonE(name),
		AddTagBsonE(tags),
		AddNegTagBsonE(negtags)};
}

func SearchRecipes(query string, db *mongo.Database) []Recipe {
	var objRecipes []Recipe
	request := QueryStringToMongoQuery(query)
	cur, err := db.Collection("Recipes").Find(context.TODO(), request)
	if err != nil {log.Fatal(err)}
	for cur.Next(context.TODO()) {
		var elem Recipe
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		elem.FillIngredients(db)
		for idx := range elem.Ingredients {
			elem.Ingredients[idx].FillUnit(db)
		}
		objRecipes = append(objRecipes, elem)
	}
	return objRecipes;
}
