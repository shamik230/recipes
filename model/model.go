package model

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	Difficulty  string   `json:"difficulty"`
	Tags        []string `json:"tags"`
}
