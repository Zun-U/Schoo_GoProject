package article

import (
	"testing"
)

func TestGetAll(t *testing.T) {

	// テストしたい関数の呼び出し
	got, err := GetAll()

	// エラーハンドリング
	if err != nil {
		t.Fatal("記事の取得に失敗しました:", err)
	}

	if len(got) != 4 {
		t.Fatal("記事数が4ではありません異なります:", len(got))
	}

	testEq(t, "自己紹介", got[0].Title)
	testEq(t, "こんなことがありました", got[1].Title)
	testEq(t, "仕事について", got[2].Title)
	testEq(t, "ブログ始めました", got[3].Title)

}

// 中身を確認するヘルパー関数の作成
func testEq(t *testing.T, want, got string) {

	// 関数の呼び出し時にエラーを検知するための関数
	t.Helper()

	if want != got {
		t.Fatalf("want: %v\ngot:  %v", want, got)
	}

}
