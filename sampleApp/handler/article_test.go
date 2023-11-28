package handler

import (
	"html/template"
	"net/http"
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
			tm := template.Must(template.ParseFS())
		})
	}

}