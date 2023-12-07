package handler

import (
	"html/template"
	"schoo/sampleApp/article"
	"schoo/sampleApp/test"
	"testing"
)

func newHandler(t *testing.T) (*Handler, func() error) {

	t.Helper()

	testdb := test.DB(t)

	a := article.New(testdb)

	h := New(
		template.Must(template.ParseFiles("../assets/index.html")),
		template.Must(template.ParseFiles("../assets/article.html")),
		a,
	)

	close := func() error {
		return testdb.Close()
	}

	return h, close

}
