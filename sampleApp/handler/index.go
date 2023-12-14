package handler

import (
	"log"
	"net/http"
	"schoo/sampleApp/article"
	"time"
)

type Summary struct {
	ID      int
	Title   string
	Summary string
	Created time.Time
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := h.article.GetAll(TableName)

	if err != nil {
		log.Println("記事一覧の取得に失敗しました")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 無名構造体
	params := struct {
		Title     string
		Summaries []Summary
	}{
		Title:     "わたしのブログ",
		Summaries: toSummaries(articles),
	}

	h.templateIndex.Execute(w, params)
}

func toSummaries(articles []article.Article) []Summary {

	summaries := make([]Summary, 0, len(articles))

	for _, v := range articles {

		summaries = append(summaries, Summary{
			ID:      v.ID,
			Title:   v.Title,
			Created: v.Created,
			Summary: summarize(v.Content, 90),
		},
		)

	}

	return summaries

}

func summarize(s string, length int) string {

	r := []rune(s) // マルチバイト対応

	if len(r) <= length {
		return s
	}

	return string(r[:length]) + "..."

}
