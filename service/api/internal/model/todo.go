package model

type TODO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type TODOs []*TODO
