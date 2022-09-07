package routes

import (
	"net/http"

	"example.com/hi/model"
	"github.com/gin-gonic/gin"
)

type ResponseRecipes struct {
	Total int            `json:"total"`
	Data  []model.Recipe `json:"data"`
}

// http

type routes struct {
	db model.Repo
}

func NewRoutes(db model.Repo) *routes {
	return &routes{db: db}
}

func (r *routes) GetRecipes(c *gin.Context) {

	recipies, err := r.db.GetRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "database error",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseRecipes{
		Total: len(recipies),
		Data:  recipies,
	})
}

func (r *routes) IndexPage(c *gin.Context) {
	recipies, err := r.db.GetRecipes()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"title": "Возникли технические шоколадки 🍫",
		})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Рецепты",
		"recepies": recipies,
		"total":    len(recipies),
	})
}
