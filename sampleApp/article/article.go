package article

import (
	"database/sql"
	"fmt"

	"time"
)

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type Service struct {
	db *sql.DB
}

func New(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAll() ([]Article, error) {

	rows, err := s.db.Query("SELECT * FROM article_test ORDER BY id DESC LIMIT 100;")

	if err != nil {
		return nil, fmt.Errorf("クエリが失敗しました: %w", err)
	}

	defer rows.Close()

	var articles []Article

	for rows.Next() {

		var article Article

		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Created); err != nil {
			return nil, fmt.Errorf("スキャン失敗しました: %w", err)
		}

		articles = append(articles, article)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *Service) Get(id int) (*Article, error) {

	// MySQLでプレイスホルダーを使用したい場合は、プリペアードステートメントを使用するを使用する
	stmt, err := s.db.Prepare("SELECT * FROM article_test WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("SQLの作成に失敗しました: %w", err)
	}
	defer stmt.Close()

	var article Article

	err = stmt.QueryRow(id).Scan(&article.ID, &article.Title, &article.Content, &article.Created)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("結果が１行も帰りませんでした: %w", err)
		}
		return nil, fmt.Errorf("スキャン失敗しました: %w", err)
	}

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	return &article, nil

}

// func articleContent() []Article {
// 	return []Article{
// 		{
// 			ID:      4,
// 			Title:   "自己紹介",
// 			Content: "今更ですが、わたしの自己紹介をさせてください。",
// 		},
// 		{
// 			ID:      3,
// 			Title:   "こんなことがありました",
// 			Content: "先日、散歩をしていたときに変な出来事に遭遇しました。",
// 		},
// 		{
// 			ID:      2,
// 			Title:   "仕事について",
// 			Content: "わたしは新卒のころからずっと続けている仕事があります。",
// 		},
// 		{
// 			ID:      1,
// 			Title:   "ブログ始めました",
// 			Content: "このブログではわたしの個人的な事柄について書くつもりです。",
// 		},
// 	}
// }
