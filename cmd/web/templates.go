package main

import "github.com/rbansalrahul6/textbin/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
