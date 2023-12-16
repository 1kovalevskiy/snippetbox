package entity

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Snippet struct {
	ID      int       `json:"id"       example:"1"`
	Title   string    `json:"title"       example:"Supper_title"`
	Content string    `json:"content"       example:"Important_content"`
	Created time.Time `json:"created"       example:""`
	Expires time.Time `json:"expires"       example:""`
}

type SnippetCreate struct {
	Title   string `json:"title"       example:"Supper_title"`
	Content string `json:"content"       example:"Important_content"`
	Expires string `json:"expires"       example:""`
}
