package lesson5

import (
	"fmt"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Lesson5Main() {
	fmt.Println("==== SQL ====")

}

func connectDB() (*sql.DB) {

	jst, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		// error handling
	}

	c := mysql.Config{
		DBName: "db",
		User: "root",
		Passwd: "password",
		Addr: "localhost:3306",
		Net: "tcp",
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
		Loc: jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())

}




