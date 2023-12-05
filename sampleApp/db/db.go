package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func New(c mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", c.FormatDSN())

	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pingに失敗しました: %w", err)
	}

	return db, err

}
