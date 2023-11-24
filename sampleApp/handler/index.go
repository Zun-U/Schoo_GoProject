package handler

import (
	"net/http"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {

	// 無名構造体
	params := struct {
		Title string
	}{
		Title: "私のブログ",
	}

	h.templateIndex.Execute(w, params)
}