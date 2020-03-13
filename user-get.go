package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "net/http"
)

type User struct {
  user_id string
  sex string
}

// ユーザ情報を取得・表示
func drawUser(w http.ResponseWriter, r *http.Request){

  // r.URL.RawQueryでクエリパラメータ取得（KunoとかmamaとかURLに打ってもらう）
  //fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)  // debug用

  // rootで作ったのでパスワードなし
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // queryにクエリパラメータ代入
  var query string
  query = "select * from user where user_id = '" + r.URL.RawQuery + "'"
  // println(query) // debug用
  rows, err := db.Query(query)

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
    fmt.Fprintf(w, "user_id : " + user.user_id)
    fmt.Fprintf(w, ", sex : " + user.sex)
  }

  err = rows.Err()
  if err != nil {
    panic(err.Error())
  }

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUser")
}

func hello(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "ユーザ情報をみるなら/user/get?***と入力してね")
  fmt.Println("Endpoint Hit: Hello")
}

func handleRequests() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/user/get", drawUser)
  // 8080ポートで起動
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests()
}
