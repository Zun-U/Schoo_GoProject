package appmain

import (
	"html/template"
	"net/http"
	"schoo/sampleApp/handler"
)

func AppMain() {

	h := handler.New(template.Must(template.ParseFiles("sampleApp/assets/index.html")), template.Must(template.ParseFiles("sampleApp/assets/article.html")))

	// ルーティング
	http.HandleFunc("/", h.Index)
	http.HandleFunc("/articles", h.Article)

	http.ListenAndServe(":8181", nil)

}
