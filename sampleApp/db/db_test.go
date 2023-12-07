package db

import (
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	// "database/sql"
)

// testのキャッシュクリア
//
// ********************
// go clean -testcache
// ********************


func TestNew(t *testing.T) {

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal("日本時間の取得に失敗しました。", err)
	}

	c := mysql.Config{
		DBName: "sample",
		User:   "gopher",
		Passwd: "password",
		// Addr:      "localhots:3306",         // DockerでDBを起動している場合は、「Localhost」で接続できない
		Addr:      "schoo_goproject-db-1:3306", // 「localhost:」ではなくDockerの「DBコンテナ名」を記載する
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := New(c)

	if err != nil {
		t.Fatal("データベースの設定に失敗しました:", err)
	}
	if db == nil {
		t.Fatal("データベースが存在しません")
	}
	_, err = db.Query("SELECT 1")
	if err != nil {
		t.Fatal("クエリに失敗しました:", err)
	}
	if err := db.Close(); err != nil {
		t.Fatal("データベース接続切断に失敗しました", err)
	}

}
