package article

type Article struct {
	Title string
}

func GetAll() ([]Article, error) {

	x := Titles()

	return x, nil
}


func Titles() []Article {
	return []Article {
		{Title:"自己紹介"},
		{Title:"こんなことがありました"},
		{Title:"仕事について"},
		{Title:"ブログ始めました"},
	}
}