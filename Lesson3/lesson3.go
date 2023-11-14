package lesson3

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

func Lesson3Main() {

	fmt.Println("==== Concurrency ====")

	// キーワード「go」でgoroutineを起動
	go print5x("o")
	go print5x("x")
	// print5x("\n")
	go print5x("\n") // goroutineはgoランタイムが管理する軽量なスレッドではあるが、並行処理の実行中にmain関数が終了してしまう場合がある

	ch := make(chan int)
	go send(1, ch)
	go send(2, ch)
	go send(3, ch)
	fmt.Println(<-ch, <-ch, <-ch)
	// <-ch 送信側も受信側も値が来るまで待っている為、送受信される値が無いとブロック(処理が停止)することになる

	// バッファ付きチャネル
	bufferChannel()

	// クローズ
	ch2 := make(chan int)
	go send3(3, ch2)
	for v := range ch2 {
		fmt.Println(v)
	}

	// セレクト
	selectChan()

	// エラーチャネル
	ch3 := make(chan string)
	resultCh, errCh := parser(ch3)

	// ss := []string{"1", "2", "3", "abc"}
	ss := []string{"1", "2", "3"}

	for _, s := range ss {
		ch3 <- s
		select {
		case result := <-resultCh:
			fmt.Printf("%d(%T)\n", result, result)
		case err := <-errCh:
			fmt.Println("error!", err)
			return
		}
	}
	close(ch3)

	// パイプライン
	rand.New(rand.NewSource(time.Now().UnixNano())) // シード値
	nums := make([]int, 0, 10)
	for i := 0; i < cap(nums); i++ {
		nums = append(nums, rand.Intn(10))
	}

	// スライスをチャネルへ変換
	fmt.Print(" x: ")
	for x := range convertChannel(nums) {
		fmt.Printf("%2d ", x)
	}
	fmt.Print("\n")

	// 2乗
	fmt.Print("*= x: ")
	for x := range sq(convertChannel(nums)) {
		fmt.Printf("%2d ", x)
	}
	fmt.Print("\n")

	// 結果を半分にする
	fmt.Print("/= 2: ")
	for x := range halve(sq(convertChannel(nums))) {
		fmt.Printf("%2d ", x)
	}
	fmt.Print("\n")

	fmt.Print("%= 7: ")
	for x := range mod7(sq(convertChannel(nums))) {
		fmt.Printf("%2d ", x)
	}
	fmt.Print("\n")
}

// 文字列を5回出力
func print5x(s string) {
	for i := 0; i < 5; i++ {
		fmt.Print(s)
		time.Sleep(200 * time.Millisecond)
	}
}

// channel
func send(x int, ch chan int) {
	ch <-x
}

// buffer
func bufferChannel() {
	ch := make(chan int, 3)
	ch <-1
	ch <-2
	ch <-3 // 受信しない
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// close
func send3(x int, ch chan int) {
	ch <- x
	ch <- x*x
	ch <- x*x*x
	close(ch)
}

// select
func gen (intCh chan int, stringCh chan string) {
	intCh <- 1
	stringCh <- "abc"
	intCh <- 2
	stringCh <- "ABC"
	intCh <- 3
	stringCh <- "done"
}

func selectChan() {

	intCh := make(chan int)
	stringCh := make(chan string)

	go gen(intCh, stringCh)

	for {
		select {
		case i := <-intCh:
			fmt.Println(i)
		case s := <-stringCh:
			fmt.Println(s)
			if s == "done" {
				return
			}
		default:
			fmt.Println(".")
		}
	}

}

// channel generator
func generator(in <-chan string) <-chan int { // 入力用のチャネルを渡す

	out := make(chan int) // 出力用のチャネルを作成

	go func() { // goroutineで入力用のチャネルに送られてきた文字列を処理
		for str := range in {
			// 何かの処理
			fmt.Println(str)
		}
	}()
	return out // 出力用のチャネルを返す
}

// error channel
func generateError(in <-chan string) (<-chan int, <-chan error) {

	out := make(chan int)
	errCh := make(chan error) //エラー用のチャネル

	go func() {
		for str := range in { //入力用のチャネルに送られてきた文字列を処理(エラー処理も行う)
			//処理
			fmt.Println(str)
		}
	}()

	return out, errCh

}

// 具体例
func parser(s <-chan string) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error)

	go func(){
		for str := range s {
			i, err := strconv.Atoi(str)
			if err != nil {
				errCh <- err
			}
			ch <- i
		}
	}()

	return ch, errCh
}

// pipeline
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x*x
		}
		close(out)
	}()
	return out
}

func halve(in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x /2
		}
		close(out)
	}()
	return out
}

func convertChannel(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			time.Sleep(200 * time.Millisecond)
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func mod7(in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x % 7
		}
		close(out)
	}()
	return out
}