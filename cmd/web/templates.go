package main

import "github.com/yeinerfonseca/snippetbox/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
