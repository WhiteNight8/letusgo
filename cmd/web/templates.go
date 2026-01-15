package main

import (
	"html/template"
	"path/filepath"

	"letgo.snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	page, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range page {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page, "./ui/html/pages/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
