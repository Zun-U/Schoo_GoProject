package article

import (
	// "fmt"
	"schoo/sampleApp/test"
	"strconv"
	"testing"
)

const TestTableName string = "article_test"

func TestGetAll(t *testing.T) {

	testdb := test.DB(t)

	defer test.Close(t, testdb)

	s := New(testdb)

	// テストしたい関数の呼び出し
	got, err := s.GetAll(TestTableName)

	// エラーハンドリング
	if err != nil {
		t.Fatal("記事の取得に失敗しました:", err)
	}

	if len(got) != 4 {
		t.Fatal("記事数が4ではありません異なります:", len(got))
	}

	test.Eq(t, "自己紹介", got[0].Title)
	test.Eq(t, "こんなことがありました", got[1].Title)
	test.Eq(t, "仕事について", got[2].Title)
	test.Eq(t, "ブログはじめました", got[3].Title)

}

// 中身を確認するヘルパー関数の作成
// func testEq(t *testing.T, want, got string) {
//
// 	// 関数の呼び出し時にエラーを検知するための関数
// 	t.Helper()
//
// 	if want != got {
// 		t.Fatalf("want: %v\ngot:  %v", want, got)
// 	}
//
// }

func TestGet(t *testing.T) {

	testdb := test.DB(t)

	defer test.Close(t, testdb)

	s := New(testdb)

	for i := 1; i < 5; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := s.Get(TestTableName, i)
			if err != nil {
				t.Fatal("記事の取得に失敗しました:", err)
			}
			if got == nil {
				t.Fatal("gotがnilです")
			}
			if got.Title == "" {
				t.Fatal("タイトルが空です")
			}
			if got.Content == "" {
				t.Fatal("中身が空です")
			}
		})
	}

	// t.Run("not found", func(t *testing.T) {
	// 	got, err := s.Get(TestTableName, -1)
	// 	if err == nil {
	// 		t.Fatal("エラーがnilです")
	// 	}
	// 	if got == nil {
	// 		t.Fatal("gotがnilです")
	// 	}
	// })
}

func TestCreateArticle(t *testing.T) {

	testdb := test.DB(t)
	// defer test.Close(t, testdb)
	defer test.Clear(t, testdb)

	s := New(testdb)

	id, err := s.Create(TestTableName, "サンプルタイトル", "これはテスト用サンプルテキストです")
	if err != nil {
		t.Fatal("failed to create article:", err)
	}

	// idのテスト
	test.Eq(t, 5, id)

	// 作成した記事を取得
	got, err := s.Get(TestTableName, id)
	if err != nil {
		t.Fatal("failed to get article:", err)
	}

	// Createに渡した内容通りか確認
	test.Eq(t, "サンプルタイトル", got.Title)
	test.Eq(t, "これはテスト用サンプルテキストです", got.Content)

}

func TestDelete(t *testing.T) {

	testdb := test.DB(t)
	defer test.Close(t, testdb)

	s := New(testdb)

	err := s.Delete(TestTableName, 1)
	if err != nil {
		t.Fatal("failed to create article:", err)
	}

	// articleを取得
	got, err := s.GetAll(TestTableName)
	if err != nil {
		t.Fatal("failed to get article:", err)
	}

	// 数が減っていればOK
	test.Eq(t, 3, len(got))

}