package main

import (
	"GoRecipes/dbops"
	"github.com/gin-gonic/gin"
	"os"
)

func config() {
	if len(os.Args) == 1 {
	}
}

func main() {
	config()
	router := gin.Default()
	router.Static("/static", "_frontend/static")
	router.LoadHTMLGlob("_frontend/html/*")

	var db = dbops.Connect("mongodb://localhost:27017")

	if dbops.DbIsEmpty(db) {
		dbops.InitWithBaseData(db, true)
	}

	mainRoutes(router, db)
	//auth_routes(router, db)

	_ = router.Run()
}
