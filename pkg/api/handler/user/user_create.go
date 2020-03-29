// 各エンドポイントの処理を定義


package HandlerPkg

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
)

/*
type User struct {
  User_id string
  Sex string
}*/

// Init
//var users []User


func HandleUserCreate(w http.ResponseWriter, r *http.Request) {

  // rootで作ったのでパスワードなし
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // MySQLにユーザ情報登録
  db.Query("insert into user values ('Kuuko', 'm');")

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUserCreate")
}
