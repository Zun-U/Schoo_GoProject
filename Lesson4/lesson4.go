package lesson4

import (
	"fmt"
	"html/template" // Goが提供する標準ライブラリ
	"log"
	"net/http" // Goが提供する標準ライブラリ
	"time"
	// "os"
	// "github.com/lib/pq"
)

func Lesson4Main() {

	fmt.Println("==== Install Go & HTML ====")

	url := setInstaller()
	url.Install()

	// 標準パッケージ 「net/http」
	// httpExample()

	// 標準パッケージ 「html/template」
	templateExample()

}

type WindowsInstallGo struct {
	URL string
}

func (wig WindowsInstallGo) Install() {
	fmt.Println(wig)
}

type Installer interface {
	install() error
}

// ファクトリ関数
func setInstaller() WindowsInstallGo {
	url := WindowsInstallGo{
		URL: "https://go.dev/dl/go1.21.4.windows-amd64.msi",
	}
	return url
}

// net/http
func httpExample() {

	x := responseStrings()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 「fmt.Fprintf」...(第一引数)に(第二引数)を書き込む
		// 「w」がレスポンス
		fmt.Fprintf(w, "Method: %s\n", r.Method)     // レスポンスを返すコード
		fmt.Fprintf(w, "Referer: %s\n", r.Referer()) // http requestのメタ情報
		fmt.Fprintf(w, "Time: %v\n", time.Now())     // 現在時刻
	})

	http.HandleFunc("/lessons", func(w http.ResponseWriter, r *http.Request){
		for _, v := range x {
			fmt.Fprintf(w, v)
				// 「http://localhost:8080/lessons」にアクセスすると以下が表示される
				// ------------------------------------------------------------
				// Go入門
				// 1.入門編
				// 2.応用編
				// 3.実践編
				// ------------------------------------------------------------
		}
	})

	// 「localhost:8080」でサーバーをListen
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// 「http://localhost:8080/」にアクセスすると以下が表示される
	// ------------------------------------------------------------
	// Method: GET
	// Referer:
	// Time: 2023-11-16 11:19:07.7985421 +0900 JST m=+31.133113701
	// ------------------------------------------------------------

}

func responseStrings() []string {

	x := []string {
		"<h1>Go入門</h1>",
		"<ol>",
		"<li>入門編</li>",
		"<li>応用編</li>",
		"<li>実践編</li>",
		"</ol>",
	}

	return x

}


type responsParams struct {
	Title 	string
	Lessons []string
	Time    time.Time // time.Time型
}

func setResponsParams() responsParams {
	params := responsParams {
		Title: "Go入門",
		Lessons: responseContents(),
		Time: time.Now(),
	}
	return params
}

// html/template
func templateExample() {

	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		log.Fatalln("index.htmlのパースに失敗しました", err)
	}

	// // 無名構造体
	// params := struct {
	// 	Title string
	// 	Lessones []string
	// }{
	// 	Title: "Go入門",
	// 	Lessones:  responseStrings(),
	// }

	params := setResponsParams()

	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request){
			t.Execute(w, params)
			// for _, v := range params.Lessons {
			// 	t.Execute(w, v)
			// }
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}


func responseContents() []string {

	x := []string {
		"入門編",
		"応用編",
		"実践編",
	}

	return x

}

