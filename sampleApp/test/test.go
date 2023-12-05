package test

import (
	"testing"
)

func Eq[T comparable](t *testing.T, want, got T) { // ジェネリクス(comparable => 比較演算子が使用できる型ならなんでも受け取れる)
	t.Helper()
	if want != got {
		t.Fatalf("\nwant: %v\ngot: %v", want, got)
	}
}