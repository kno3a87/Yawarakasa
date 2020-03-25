// 各エンドポイントの処理を定義


package HandlerPkg

import (
  "context"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
  // json定義
  w.Header().Set("Content-Type", "application/json")

  // rootで作ったのでパスワードなし
  db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/YawarakasaApp")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  url := ""

  request, err := http.NewRequest("GET", url, nil)
  
  // リクエストヘッダに代入
  request.Header.Set("userId", "Kuno")
  if err != nil{
    //log.Fatal(err)
  }

  // リクエストヘッダから取得してcontextに保存
  // 空のコンテキスト作成
  ctx := context.Background()
  // コンテキストに代入
  ctx = context.WithValue(ctx, "userId", request.Header.Get("userId"))
  // Terminalにログ表示
  fmt.Println(request.Header.Get("userId"))
  fmt.Println(ctx.Value("userId"))

  //クエリパラメータ
  /*
  params := request.URL.Query()
  params.Add("userId","Kuno") // 直接指定してみる
  request.URL.RawQuery = params.Encode()
  */

  // Terminalにログ表示
  /*
  fmt.Println(params) // 'map[userId:[Kuno]]'
  fmt.Println(params["userId"]) // '[Kuno]'
  fmt.Println(request.URL.String())  // '?userId=Kuno'
  */


  // MySQLからuserId=Kunoの情報取ってくる
  rows, err := db.Query("select * from user where user_id = ?", ctx.Value("userId"))
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
  fmt.Println("Endpoint Hit: drawUser")
}
