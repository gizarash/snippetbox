package main

import "github.com/gizarash/snippetbox/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}