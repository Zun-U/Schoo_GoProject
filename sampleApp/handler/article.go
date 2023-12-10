package handler

import (
	"fmt"
	"log"
	"net/http"
	"schoo/sampleApp/article"
	"strconv"
	"unicode/utf8"
)

type ArticleContent struct {
	Title   string
	Article *article.Article
}

const TableName = "article_test"

// func (h *Handler) Article(w http.ResponseWriter, req *http.Request) {
// 	queryID := req.URL.Query().Get("id")
// 	id, err := strconv.Atoi(queryID)
//
// 	if err != nil {
// 		log.Printf("クエリパラメータのパースに失敗しました: %s", queryID)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
//
// 	a, err := h.article.Get(TableName, id)
// 	if err != nil {
// 		log.Printf("記事の取得に失敗しました: %d", err)
// 		w.WriteHeader(http.StatusNotFound)
// 	}
//
// 	params := ArticleContent{
// 		Title:   "私のブログ",
// 		Article: a,
// 	}
//
// 	h.templateArticle.Execute(w, params)
//
// }

type IndexTitle struct {
	Title string
}

func (h *Handler) NewArticle(w http.ResponseWriter, r *http.Request) {

	params := IndexTitle{
		Title: "わたしのブログ",
	}

	h.templateNewArticle.Execute(w, params)

}

func (h *Handler) Article(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.getArticle(w, r)
	case http.MethodPost:
		h.createArticle(w, r)
	case http.MethodDelete:
		h.deleteArticle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *Handler) getArticle(w http.ResponseWriter, r *http.Request) {

	queryID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)

	if err != nil {
		log.Printf("クエリパラメータのパースに失敗しました: %s", queryID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, err := h.article.Get(TableName, id)
	if err != nil {
		log.Printf("記事の取得に失敗しました: %d", err)
		w.WriteHeader(http.StatusNotFound)
	}

	params := ArticleContent{
		Title:   "わたしのブログ",
		Article: a,
	}

	h.templateArticle.Execute(w, params)

}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {

	// POSTされてきた値の取得
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	// バリデーション
	if err := validate(title, content); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// 投稿
	id, err := h.article.Create(TableName, title, content)
	if err != nil {
		log.Println("記事の作成に失敗しました", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", fmt.Sprintf("/articles?id=%d", id))
	w.WriteHeader(http.StatusSeeOther)

}

func validate(title, content string) error {

	// 長さチェック
	if len(title) == 0 {
		return fmt.Errorf("タイトルを入力してください")
	}
	if len(content) == 0 {
		return fmt.Errorf("内容を入力してください")
	}

	if utf8.RuneCountInString(title) > 100 {
		return fmt.Errorf("タイトルは100文字以内に収めてください")
	}

	if utf8.RuneCountInString(content) > 1000 {
		return fmt.Errorf("内容は1000文字以内に収めてください")
	}

	return nil
}

func (h *Handler) deleteArticle(w http.ResponseWriter, r *http.Request) {

	queryID := r.URL.Query().Get("id")

	id, err := strconv.Atoi(queryID)
	if err != nil {
		log.Println("failed to parse query:", queryID, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.article.Delete(id); err != nil {
		log.Println("failed to delete article:", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
