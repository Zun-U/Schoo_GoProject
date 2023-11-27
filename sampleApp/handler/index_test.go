package handler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
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

	// 記事一覧のテスト
	titles := []string{
		"自己紹介",
		"こんなことがありました",
		"仕事について",
		"ブログ始めました",
	}

	for _, title := range titles {
		if !strings.Contains(rec.Body.String(), fmt.Sprintf("<h2>%s</h2>", title)) {
			t.Fatal("article title is missing", title)
		}
	}


}
