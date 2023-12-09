package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testTable struct {
	Name string
	ID   string
	Want int
}

// テーブル駆動テスト
func TestArticle(t *testing.T) {

	// テーブル
	tests := []testTable{
		{
			Name: "success",
			ID:   "1",
			Want: http.StatusOK,
		},
		{
			Name: "bad request",
			ID:   "abc",
			Want: http.StatusBadRequest,
		},
		{
			Name: "not found",
			ID:   "-1",
			Want: http.StatusNotFound,
		},
	}

	h, close := newHandler(t)
	defer close()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/articles?id="+tt.ID, nil)
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			if rec.Code != tt.Want {
				t.Fatalf("want: %d, got: %d", tt.Want, rec.Code)
			}

			if rec.Code == http.StatusOK {
				// 記事のタイトルチェック
				if !strings.Contains(rec.Body.String(), "<h2>") {
					t.Fatal("記事のタイトルがありません")
				}

				// 記事の中身チェック
				if !strings.Contains(rec.Body.String(), "<p>") {
					t.Fatal("記事の内容がありません")
				}
			}
		})
	}

}

func TestNewArticle(t *testing.T) {

	h, close := newHandler(t)
	defer close()

	req := httptest.NewRequest(http.MethodGet, "/articles/new", nil)
	rec := httptest.NewRecorder()
	h.NewArticle(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("status code is not 200:", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "<h1>わたしのブログ</h1>") {
		t.Fatal("タイトルがありません")
	}

	if !strings.Contains(rec.Body.String(), "<h2>新規作成</h2>") {
		t.Fatal("サブタイトルがありません")
	}

	if !strings.Contains(rec.Body.String(), `<form action="/articles" method="post">`) {
		t.Fatal("フォームがありません")
	}

}
