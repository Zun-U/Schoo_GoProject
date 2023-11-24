package appmain

import (
	"html/template"
	"net/http"
	"schoo/sampleApp/handler"
)

func AppMain() {

	h := handler.New(template.Must(template.ParseFiles("sampleApp/assets/index.html")))

	http.HandleFunc("/", h.Index)

	http.ListenAndServe(":8181", nil)

}
