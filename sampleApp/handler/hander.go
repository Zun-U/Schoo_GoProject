package handler

import (
	"html/template"
	"schoo/sampleApp/article"
)

type Handler struct {
	templateIndex      *template.Template
	templateArticle    *template.Template
	templateNewArticle *template.Template
	article            *article.Service
}

func New(templateIndex, templateArticle, templateNewArticle *template.Template, article *article.Service) *Handler {
	return &Handler{
		templateIndex:      templateIndex,
		templateArticle:    templateArticle,
		templateNewArticle: templateNewArticle,
		article:            article,
	}
}
