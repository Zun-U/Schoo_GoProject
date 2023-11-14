package lesson2

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"os"
	"strconv"
)

func Lesson2Main() {

	fmt.Println("==== Function and Method ====")

	fmt.Println(cube(2))
	fmt.Println(cube(5))

	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))

	s := "abcdefg"
	fmt.Println(split(s))

	fmt.Println(max(1, 2, 3, 4))

	fmt.Println(mul(2, 3))
	fmt.Println(div(3, 4))

	pointerExample()

	me := Profile{
		Name:  "鈴木",
		Email: "test@example.com",
		Age:   20,
	}

	fmt.Println(me)
	fmt.Println(me.Age)

	a := Square{
		Width:  4,
		Height: 3,
	}

	b := Square{
		Width:  8,
		Height: 3,
	}

	fmt.Println(a.Area(), b.Area())

	fmt.Println(a.Diagonal(), b.Diagonal())

	var e, f, g Age = 16, 18, 20
	fmt.Println(e, e.IsAdult())
	fmt.Println(f, f.IsAdult())
	fmt.Println(g, g.IsAdult())

	ss := Square{
		Width:  20,
		Height: 40,
	}
	ss.Double()
	fmt.Println(s)
	ss.Halve()
	fmt.Println(ss)

	ss.Swap()
	fmt.Println(ss)

	o := Square{
		Width:  2,
		Height: 3,
	}

	t := Triangle{
		Base:   3,
		Height: 4,
	}

	Shock(o)
	Shock(t)

	Shock(Circle{1})

	errorHandlingExample()

	check("10", "9", "19", "333", "abc")

	defer fmt.Println("defer実行")

}

// 関数
func cube(x int) int {
	return x * x * x
}

func add(x int, y int) int {
	return x + y
}

// 引数の型省略
func sub(x, y int) int {
	return x - y
}

// 複数の戻り値
func split(s string) (string, string) {
	l := len(s) / 2
	return s[:l], s[l:]
}

// 可変長引数
func max(x ...int) int {
	switch len(x) {
	case 1:
		return x[0]
	case 2:
		if x[0] > x[1] {
			return x[0]
		}
	default:
		return x[1]
	}
	return max(x[0], max(x[1:]...))
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) (float32, error) {
	if y == 0 {
		return 0, errors.New("0での除算は出来ません")
	}
	return float32(x) / float32(y), nil
}

// ポインタ
func pointerExample() {
	x := 10
	p := &x
	fmt.Println(p, *p)

	// 参照元の値を書き換え
	*p = 20
	fmt.Println(x)

	// 値の再代入で参照側も変わる
	x = 30
	fmt.Println(*p)
}

// 構造体
type Profile struct {
	Name  string
	Email string
	Age   int
}

type Square struct {
	Width  float64
	Height float64
}

// メソッド
func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) Diagonal() float64 {
	return math.Sqrt(s.Width*s.Width + s.Height*s.Height)
}

// ユーザー定義型
type Age int

func (a Age) IsAdult() bool {
	return a >= 18
}

func (a Age) String() string {
	if a.IsAdult() {
		return "成人"
	}
	return "未成年"
}

func (s Square) Double() {
	s.Width *= 2
	s.Height *= 2
}

// 構造体のポインタ
func (s *Square) Halve() {
	s.Width /= 2
	s.Height /= 2
}

func (s *Square) Swap() {
	s.Width, s.Height = s.Height, s.Width
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// インターフェース
type Polygon interface {
	Area() float64
}

func Shock(p Polygon) {
	fmt.Printf("%T%vの面積は%fです\n", p, p, p.Area())
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

// error
func errorHandlingExample() {

	// Goの標準なエラーハンドリング
	if _, err := strconv.Atoi("abc"); err != nil {
		fmt.Println(err)
	}

	if _, err := url.Parse("http://"); err != nil {
		fmt.Println(err)
	}

	if _, err := os.Open("a.txt"); err != nil {
		fmt.Println(err)
	}

}

func atoi(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("関数atoiで失敗しました: %w", err)
	}
	fmt.Println(s, i)
	return nil
}

func check(values ...string) {
	for _, s := range values {
		if err := atoi(s); err != nil {
			fmt.Println("checkエラー:", err)
		}
	}
}
