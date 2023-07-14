package main

import "github.com/zacharysyoung/SnippetBox/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
