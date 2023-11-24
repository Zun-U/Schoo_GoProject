package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Index(t *testing.T) {
	tm := template.Must(template.ParseFiles("../assets/index.html"))
	h := New(tm)

	// Topページ(/)を取得(Get)
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()

	h.Index(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("ステータスコードが200ではありません。", rec.Code)
	}

	// トップページのタイトルのテスト
	if !strings.Contains(rec.Body.String(), "<h1>私のブログ</h1>") {
		t.Fatal("タイトルがありません。")
	}
}
