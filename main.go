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

// Anti Virus SoftによってGoが実行できない場合
//
// 「Norton」の場合
// 「設定」→「スキャンとリスク」で「自動保護....検出から除外する項目」で「AppData\Local」以下の「staticcheck」、「gopls」、「tempファイル」を指定する。