package main

import (
	"schoo/Lesson1"
	"schoo/Lesson2"
)

func main() {

	// 変数、定数、制御構造、配列
	lesson1.Lesson1Main()

	// 関数とメソッド
	lesson2.Lesson2Main()

}

// アンチウィルスによってGoが実行できない場合
//
// 「Norton」の場合
// 「設定」→「スキャンとリスク」で「自動保護....検出から除外する項目」で「AppData\Local」以下の「staticcheck」、「gopls」、「temp」ファイルを指定する。
//
// 参考URL：
// https://stackoverflow.com/questions/43019581/go-lang-access-denied
// https://community.norton.com/en/forums/heuradvmlb-has-been-detected-false-positive