package main

import (
	"fmt"

	"example.com/hi/repo"
	"example.com/hi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repo.NewRepo("postgres://sunny:1@localhost:5432/gxg")
	if err != nil {
		fmt.Println(err)
		return
	}

	routes := routes.NewRoutes(db)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/recipe", routes.GetRecipes)
	router.GET("/", routes.IndexPage)
	router.Run()
}
