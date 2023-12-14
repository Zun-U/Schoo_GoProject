package appmain

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	// "os"
	"schoo/sampleApp/article"
	"schoo/sampleApp/db"
	"schoo/sampleApp/handler"
	"time"

	"github.com/go-sql-driver/mysql"
)

// ====================
// ビルド実行時にHTMLファイルも一緒にバイナリ形式にコンパイルする
// (コメント行も必要)

//go:embed assets
var assets embed.FS

// ====================

func AppMain() {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := mysql.Config{
		DBName:    "sample",
		User:      "gopher",
		Passwd:    "password",
		Addr:      "schoo_goproject-db-1:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	d, err := db.New(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	a := article.New(d)

	h := handler.New(
		// template.Must(template.ParseFiles("sampleApp/assets/index.html")),
		// template.Must(template.ParseFiles("sampleApp/assets/article.html")),
		// template.Must(template.ParseFiles("sampleApp/assets/new.html")),
		template.Must(template.ParseFS(assets, "assets/index.html")),    // HTMLファイルをビルドファイルに含める為、「ParseFS」
		template.Must(template.ParseFS(assets, "assets/article.html")),  // HTMLファイルをビルドファイルに含める為、「ParseFS」
		template.Must(template.ParseFS(assets, "assets/new.html")),      // HTMLファイルをビルドファイルに含める為、「ParseFS」
		a,
	)

	// ルーティング
	http.HandleFunc("/", h.Index)
	http.HandleFunc("/articles", h.Article)
	http.HandleFunc("/articles/new", h.NewArticle)

	http.ListenAndServe(":8181", nil)

}

// func getEnv(key, defaultValue string) string {
//
// 	// os.Getenv 環境変数の取得(本番環境はこちらを使用する)
// 	if env := os.Getenv(key); env != "" {
// 		return env
// 	}
// 	return defaultValue
// }
