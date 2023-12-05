package handler

import (
	"log"
	"net/http"
	"schoo/sampleApp/article"
	"strconv"
)

type ArticleContent struct {
	Title   string
	Article *article.Article
}

func (h *Handler) Article(w http.ResponseWriter, req *http.Request) {
	queryID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)

	if err != nil {
		log.Printf("クエリパラメータのパースに失敗しました: %s", queryID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, err := article.Get(id)
	if err != nil {
		log.Printf("記事の取得に失敗しました: %d", err)
		w.WriteHeader(http.StatusNotFound)
	}

	params := ArticleContent{
		Title:   "私のブログ",
		Article: a,
	}

	h.templateArticle.Execute(w, params)

}
