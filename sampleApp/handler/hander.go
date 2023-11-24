package handler

import (
	"html/template"
)

type Handler struct {
	templateIndex *template.Template
}

func New(templateIndex *template.Template) *Handler {
	return &Handler{
		templateIndex: templateIndex,
	}
}
