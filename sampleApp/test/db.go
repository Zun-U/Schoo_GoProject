package test

import (
	"database/sql"
	"schoo/sampleApp/db"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
)

func DB(t *testing.T) *sql.DB {

	t.Helper()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal("日本時間の取得に失敗しました。", err)
	}

	c := mysql.Config{
		DBName:    "sample",
		User:      "gopher",
		Passwd:    "password",
		Addr:      "schoo_goproject-db-1:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	d, err := db.New(c)

	if err != nil {
		t.Fatal(err)
	}

	Clear(t, d)

	// _, err = d.Exec(`
	// 	INSERT INTO article_test (id, title, content, created)
	// 	VALUES
	// 	 	(1, 'ブログはじめました', 'このブログではわたしの個人的な事柄について書くつもりです。', now()),
	// 		(2, '仕事について', 'わたしは新卒のころからずっと続けている仕事があります。', now()),
	// 		(3, 'こんなことがありました', '先日、散歩をしていた時に変な出来事に遭遇しました。', now()),
	// 		(4, '自己紹介', 'わたしの自己紹介をさせてください。', now())
	// `)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	return d

}

func Clear(t *testing.T, db *sql.DB) {
	t.Helper()
	if _, err := db.Exec("TRUNCATE TABLE article_test;"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("ALTER TABLE article_test AUTO_INCREMENT = 1;"); err != nil {
		t.Fatal(err)
	}
}

func Close(t *testing.T, db *sql.DB) {
	t.Helper()
	Clear(t, db)
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
