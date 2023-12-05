package handler

import (
	"log"
	"net/http"
	"schoo/sampleApp/article"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := article.GetAll()

	if err != nil {
		log.Println("記事一覧の取得に失敗しました")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 無名構造体
	params := struct {
		Title    string
		Articles []article.Article
	}{
		Title:    "私のブログ",
		Articles: articles,
	}

	h.templateIndex.Execute(w, params)
}
