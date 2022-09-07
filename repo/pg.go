package repo

import (
	"context"
	"fmt"

	"example.com/hi/model"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Database
type repo struct {
	db *pgxpool.Pool
}

func NewRepo(url string) (*repo, error) {
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return &repo{
		db: pool,
	}, nil
}

func (r *repo) GetRecipes() ([]model.Recipe, error) {
	rows, _ := r.db.Query(
		context.Background(),
		`select id, name, description, dificulty from recipies`,
	)

	if rows.Err() == pgx.ErrNoRows {
		return nil, nil
	}

	result := []model.Recipe{}
	for rows.Next() {
		recipe := model.Recipe{}
		err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Difficulty)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		recipe.Tags, err = r.getTags(recipe.ID)
		if err != nil {
			err = fmt.Errorf("repo.GetTags: %w", err)
			return nil, err
		}

		result = append(result, recipe)
	}

	return result, nil
}

func (r *repo) getTags(id int) (tags []string, err error) {
	rows, _ := r.db.Query(
		context.Background(),
		`select tag from recipies_tags where id = $1`,
		id,
	)

	if rows.Err() == pgx.ErrNoRows {
		return
	}

	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			err = fmt.Errorf("scan error: %w", err)
			return
		}

		tags = append(tags, tag)
	}

	return
}
