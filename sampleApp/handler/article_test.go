package handler

import (
	"html/template"
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
func TestArticle(t *testing.T){

	// テーブル
	tests := []testTable{
		{
			Name:"success",
			ID:  "1",
			Want:http.StatusOK,
		},
		{
			Name: "bad request",
			ID: "abc",
			Want: http.StatusBadRequest,
		},
		{
			Name: "not found",
			ID: "-1",
			Want: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func (t *testing.T)  {
			tm := template.Must(template.ParseFiles("../assets/article.html"))
			h := New(nil, tm)

			req := httptest.NewRequest(http.MethodGet, "/articles?id="+tt.ID, nil)
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			if rec.Code != tt.Want {
				t.Fatalf("want: %d, got: %d", tt.Want, rec.Code)
			}

			if rec.Code == http.StatusOK {
				// 記事のタイトルチェック
				if !strings.Contains(rec.Body.String(),"<h2>") {
					t.Fatal("記事のタイトルがありません")
				}

				// 記事の中身チェック
				if !strings.Contains(rec.Body.String(),"<p>") {
					t.Fatal("記事の内容がありません")
				}
			}
		})
	}

}