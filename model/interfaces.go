package model

type Repo interface {
	GetRecipes() ([]Recipe, error)
}
