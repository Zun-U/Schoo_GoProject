package appmain

import (
	"html/template"
	"net/http"
	"schoo/sampleApp/article"
)

func AppMain() {

	h := handler.New(template.Must(template.ParaseFiles("sampleApp/assets/index.html")))

	http.HandleFunc("/", h.index)

	http.ListenAndServe(":8080", nil)

}