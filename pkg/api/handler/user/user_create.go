// 各エンドポイントの処理を定義


package HandlerPkg

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
)

func HandleUserCreate(w http.ResponseWriter, r *http.Request) {

  // rootで作ったのでパスワードなし
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  url := ""
  request, err := http.NewRequest("GET", url, nil)
  // リクエストヘッダに代入
  request.Header.Set("userId", "baba")
  request.Header.Add("sex", "f")

  // MySQLにユーザ情報登録
  //db.Query("insert into user values ('Kuuko', 'm');") // debug
  db.Query("insert into user values (?, ?);", request.Header.Get("userId"), request.Header.Get("sex"))

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUserCreate")
}
