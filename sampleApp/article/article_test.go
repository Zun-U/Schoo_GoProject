package article

import (
	// "fmt"
	"schoo/sampleApp/test"
	"strconv"
	"testing"
)

func TestGetAll(t *testing.T) {

	testdb := test.DB(t)

	defer test.Close(t, testdb)

	s := New(testdb)

	// テストしたい関数の呼び出し
	got, err := s.GetAll()

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
			got, err := s.Get(i)
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

	t.Run("not found", func(t *testing.T) {
		got, err := s.Get(-1)
		if err == nil {
			t.Fatal("エラーがnilです")
		}
		if got == nil {
			t.Fatal("gotがnilです")
		}
	})
}
