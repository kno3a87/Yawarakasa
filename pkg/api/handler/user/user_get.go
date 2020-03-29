// 各エンドポイントの処理を定義


package HandlerPkg

import (
  "encoding/json"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
)

type User struct {
  User_id string `json:"ユーザID"`
  Sex string `json:"性別"`
}

// Init
var users []User


func HandleUserGet(w http.ResponseWriter, r *http.Request) {
  // context.Context受け取り
  Ctx := r.Context()
  fmt.Println(Ctx.Value("userId"))  // debug

  // json定義
  w.Header().Set("Content-Type", "application/json")

  // rootで作ったのでパスワードなし
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  // MySQLからuserId=Kunoの情報取ってくる
  //rows, err := db.Query("select * from user where user_id = ?", Ctx.Value("userId"))
  rows, err := db.Query("select * from user where user_id = 'baba'") // debug
  for rows.Next() {
    var user User
    err := rows.Scan(&user.User_id, &user.Sex)

    // jsonエンコード
    outputJson, err := json.Marshal(user)
    if err != nil {
      panic(err)
    }
    //fmt.Fprintln(User)
    fmt.Fprint(w, string(outputJson))

    if err != nil {
      panic(err.Error())
      return
    }
    //fmt.Println(user)
  }

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUserGet")
}
