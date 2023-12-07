package handler

import (
	"html/template"
	"schoo/sampleApp/article"
)

type Handler struct {
	templateIndex   *template.Template
	templateArticle *template.Template
	article         *article.Service
}

func New(templateIndex, templateArticle *template.Template, article *article.Service) *Handler {
	return &Handler{
		templateIndex:   templateIndex,
		templateArticle: templateArticle,
		article:         article,
	}
}
