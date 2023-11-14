package lesson1

import (
	"fmt"
	"math/rand"
	"time"
)

func Lesson1Main() {

	fmt.Println("==== Variable, Constant, Control Structure, Array, Slice, Map ====")

	// import
	fmt.Println(time.Now())

	// 変数の宣言
	var x int
	var y int

	fmt.Println("x:", x, " y:", y)

	// 値の代入
	x = 1
	y = 2

	fmt.Println(x, y)

	// 演算
	x = x + y
	fmt.Println(x, y)

	// 定数
	const v = 10
	fmt.Println(v)

	// 文字列型、浮動小数点型、真偽型
	var s string = "abc"
	var f float64 = 3.24
	var b bool = true

	fmt.Println(s, f, b)

	// キャスト (浮動小数点型 → 整数型)
	var z int = int(f)
	fmt.Println(z)

	// キャスト (整数型 → 浮動小数点型)
	var ff float32 = float32(v) // 整数型に変換される際に小数点以下は切り捨て
	fmt.Println(ff)

	// 型の省略(コンパイラがリテラルの型を自動で判断する)
	var xx = 100

	// 関数内で利用できる、もう一つの変数の宣言(型を省略)
	zz := 1000
	fmt.Println(xx, zz)

	// if文
	// ※ifブロックないでのみ有効な変数の宣言が行える
	if xxx := 10; (x - 2) == 0 {
		fmt.Println(xxx - 2)
	} else if (xxx - 10) == 0 {
		fmt.Println(xxx - 10)
	} else {
		fmt.Println(xxx)
	}

	// for文は以下で成り立つ
	// 1.「初期設定」
	// 2.「条件」
	// 3.「再設定」
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 10
	// 「条件」のみのfor文
	for sum < 1000 {
		fmt.Println(sum)
		sum *= sum
		fmt.Println(sum)
	}

	// FizzBuzz
	for i := 1; i < 21; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
			continue // continueを入れた方が、ネストされたif文よりわかりやすい
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	rand.New(rand.NewSource(time.Now().UnixNano())) // シード値

	// switch文
	// ※switch文もブロック内のみで有効な変数を宣言できる。(「n := rand.Intn(10);」の部分)
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("少し小さすぎます", n)
	case n > 5:
		fmt.Println("大きすぎます", n)
	default:
		fmt.Println("良い感じの数字です", n)
	}

	// 配列の宣言
	arr := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(arr)
	fmt.Println(arr[2])

	// for-range文
	// インデックスと値の取り出し
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 配列のスライス
	s1 := arr[1:3]
	s2 := arr[1:5]
	s3 := arr[4:]

	fmt.Println(s1, s2, s3)

	s1[0], s2[0] = 0, 0
	fmt.Println(s1, s2)

	// スライスの宣言
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice1))

	// append
	// スライスに要素を追加する
	slice1 = append(slice1, 17, 19)
	fmt.Println(slice1, len(slice1))

	for i, v := range slice1 {
		slice1[i] = v * 2
	}

	fmt.Println(slice1)

	// マップ
	m := map[string]int{
		"佐藤": 100,
		"鈴木": 90,
		"田中": 95,
	}

	// キーと値の取り出し
	for k, v := range m {
		fmt.Println(k, v)
	}

}
