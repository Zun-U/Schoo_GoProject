package article

import (
	"fmt"
)

type Article struct {
	Title string
}

func GetAll() ([]Article, error) {
	return []Article, nil
}