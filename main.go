package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("This is Sample Schoo Project")

	fmt.Println(time.Now())

	var x int
	var y int

	fmt.Println("x:", x, " y:", y)

	x = 1
	y = 2

	fmt.Println(x, y)

	x = x + y
	fmt.Println(x, y)

	const v = 10
	fmt.Println(v)

	var s string = "abc"
	var f float64 = 3.24
	var b bool = true

	fmt.Println(s, f, b)

	var z int = int(f)
	fmt.Println(z)

	var ff float32 = float32(v)
	fmt.Println(ff)

	var xx = 100
	zz := 1000
	fmt.Println(xx, zz)

	if xxx := 10; (x - 2) == 0 {
		fmt.Println(xxx - 2)
	} else if (xxx - 10) == 0 {
		fmt.Println(xxx - 10)
	} else {
		fmt.Println(xxx)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 10;
	for sum < 1000 {
		fmt.Println(sum)
		sum *= sum
		fmt.Println(sum)
	}

	// FizzBuzz
	for i := 1; i < 21; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
			continue
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

	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("少し小さすぎます", n)
	case n > 5:
		fmt.Println("大きすぎます", n)
	default:
		fmt.Println("良い感じの数字です", n)
	}

	arr := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(arr)
	fmt.Println(arr[2])

	for i, v := range arr {
		fmt.Println(i, v)
	}

	s1 := arr[1:3]
	s2 := arr[1:5]
	s3 := arr[4:]

	fmt.Println(s1, s2, s3)

	s1[0], s2[0] = 0, 0
	fmt.Println(s1[0], s2[0])

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice1))

	slice1 = append(slice1, 17, 19)
	fmt.Println(slice1, len(slice1))

	for i, v := range slice1 {
		slice1[i] = v * 2
	}

	fmt.Println(slice1)


	m := map[string]int{
		"佐藤":100,
		"鈴木":90,
		"田中":95,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}

// 1.
//
// 大量アクセスに強い(並行処理が強いGoroutin)
