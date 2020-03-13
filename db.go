package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	user_id   string
	sex string
}

func main() {
  // rootで作ったのでパスワードなし
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

  // user情報全部表示
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.user_id, &user.sex)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.user_id, user.sex)
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}
