package handler

import (
	"html/template"
)

type Handler struct {
	templateIndex   *template.Template
	templateArticle *template.Template
}

func New(templateIndex, templateArticle *template.Template) *Handler {
	return &Handler{
		templateIndex:   templateIndex,
		templateArticle: templateArticle,
	}
}
