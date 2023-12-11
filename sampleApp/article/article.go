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

func (s *Service) GetAll(table string) ([]Article, error) {

	rows, err := s.db.Query("SELECT * FROM " + table + " ORDER BY id DESC LIMIT 100;")

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

func (s *Service) Get(table string, id int) (*Article, error) {

	// MySQLでプレイスホルダーを使用したい場合は、プリペアードステートメントを使用するを使用することもできる
	stmt, err := s.db.Prepare("SELECT * FROM " + table + " WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("SQLの作成に失敗しました: %w", err)
	}
	defer stmt.Close()

	var article Article

	err = stmt.QueryRow(id).Scan(&article.ID, &article.Title, &article.Content, &article.Created)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("結果行が存在しません: %w", err)
		}
		return nil, fmt.Errorf("スキャン失敗しました: %w", err)
	}

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	return &article, nil

}

func (s *Service) Create(table, title, content string) (int, error) {

	stmt, err := s.db.Prepare("INSERT INTO " + table + " (title, content, created) VALUES(?, ?, now())")
	if err != nil {
		return 0, fmt.Errorf("SQLの作成に失敗しました: %w", err)
	}
	defer stmt.Close()

	var id int

	result, err := stmt.Exec(title, content)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("クエリに失敗しました: %w", err)
		}
		return 0, fmt.Errorf("スキャンに失敗しました: %w", err)
	}

	i, err := result.LastInsertId()
	if err != nil {
		fmt.Println(i)
		return 0, fmt.Errorf("インサートに失敗しました1: %w", err)
	}
	fmt.Printf("LastInsertId： %d\n", id)

	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println(num)
		return 0, fmt.Errorf("インサートに失敗しました2: %w", err)
	}

	// if err := row.Scan(&id); err != nil {
	// 	return 0, fmt.Errorf("インサートに失敗しました1: %w", err)
	// }

	// if err := row.Err(); err != nil {
	// 	return 0, fmt.Errorf("インサートに失敗しました2: %w", err)
	// }

	id = int(i)

	return id, nil

}

func (s *Service) Delete(table string, id int) error {

	_, err := s.db.Exec(`DELETE FROM `+table+` WHERE id = ?`, id)

	if err != nil {
		return fmt.Errorf("削除に失敗しました: %w", err)
	}

	return nil

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
