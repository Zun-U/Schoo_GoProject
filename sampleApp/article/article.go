package article

import (
	"fmt"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

func Get(id int) (*Article, error) {

	articles := articleContent()

	for _, v := range articles {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("記事が見つかりません: id: %d", id)
}

func GetAll() ([]Article, error) {

	x := articleContent()

	return x, nil
}

func articleContent() []Article {
	return []Article{
		{
			ID: 4,
			Title: "自己紹介",
			Content: "今更ですが、わたしの自己紹介をさせてください。",
		},
		{
			ID: 3,
			Title: "こんなことがありました",
			Content: "先日、散歩をしていたときに変な出来事に遭遇しました。",
		},
		{
			ID: 2,
			Title: "仕事について",
			Content: "わたしは新卒のころからずっと続けている仕事があります。",
		},
		{
			ID: 1,
			Title: "ブログ始めました",
			Content: "このブログではわたしの個人的な事柄について書くつもりです。",
		},
	}
}