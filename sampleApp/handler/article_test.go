package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"schoo/sampleApp/test"
	"strings"
	"testing"
)

type testTable struct {
	Name string
	ID   string
	Want int
}

// テーブル駆動テスト
func TestGetArticle(t *testing.T) {

	// テーブル
	tests := []testTable{
		{
			Name: "success",
			ID:   "1",
			Want: http.StatusOK,
		},
		// {
		// 	Name: "bad request",
		// 	ID:   "abc",
		// 	Want: http.StatusBadRequest,
		// },
		// {
		// 	Name: "not found",
		// 	ID:   "-1",
		// 	Want: http.StatusNotFound,
		// },
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
				if !strings.Contains(rec.Body.String(), "<h2 class=\"text-start py-1\">") {
					t.Fatal("記事のタイトルがありません")
				}

				// 記事の中身チェック
				if !strings.Contains(rec.Body.String(), "<p class=\"my-4\">") {
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

	if !strings.Contains(rec.Body.String(), "<h1 class=\"fs-2\">わたしのブログ</h1>") {
		t.Fatal("タイトルがありません")
	}

	if !strings.Contains(rec.Body.String(), "<h2 class=\"fs-2\">新規作成</h2>") {
		t.Fatal("サブタイトルがありません")
	}

	if !strings.Contains(rec.Body.String(), `<form action="/articles" method="post" class="my-3">`) {
		t.Fatal("フォームがありません")
	}

}

// テーブル駆動テスト用のテストテーブル構造体
type createAritcleTestTable struct {
	name    string
	title   string
	content string
	want    int
}

func TestCreateArticle(t *testing.T) {

	h, close := newHandler(t)
	defer close()

	// テスト用の値を作成(カラムの許容数を「title => 100」「content => 1000」にしておくこと)
	tooLongTitle   := strings.Repeat("あ", 101)
	tooLongContent := strings.Repeat("い", 1001)
	validTitle     := strings.Repeat("う", 100)
	validContent   := strings.Repeat("え", 1000)

	// テーブル作成
	tests := []createAritcleTestTable{
		{
			name: "too long title",
			title: tooLongTitle,
			content: validContent,
			want: http.StatusBadRequest,
		},
		{
			name: "too long content",
			title: validTitle,
			content: tooLongContent,
			want: http.StatusBadRequest,
		},
		{
			name: "success",
			title: validTitle,
			content: validContent,
			want: http.StatusSeeOther,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			// リクエストボディ
			form := url.Values {
				"title":   {tt.title},
				"content": {tt.content},
			}
			body := strings.NewReader(form.Encode())

			// リクエストの作成
			req := httptest.NewRequest(http.MethodPost, "/articles", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			// 実行
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			// ステータスコードのテスト
			test.Eq(t, tt.want, rec.Code)

			// レスポンスのテスト
			if tt.want == http.StatusSeeOther {
				test.Eq(t, "/articles?id=5", rec.Header().Get("Location"))
			}
		})
	}

}


type deleteArticleTestTable struct {
	name string
	id   string
	want int
}

func TestDeleteArticle(t *testing.T) {

	tests := []deleteArticleTestTable{
		{
			name: "success",
			id: "1",
			want: http.StatusNoContent,
		},
		{
			name: "bad request",
			id: "abc",
			want: http.StatusBadRequest,
		},
		{
			name: "do nothing",
			id: "-1",
			want: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h, close := newHandler(t)
			defer close()

			req := httptest.NewRequest(http.MethodDelete, "/articles?id=" + tt.id, nil)
			rec := httptest.NewRecorder()

			h.Article(rec, req)

			test.Eq(t, tt.want, rec.Code)

		})
	}

}