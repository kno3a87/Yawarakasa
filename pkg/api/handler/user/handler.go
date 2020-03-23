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
  user_id string `json:"ユーザID"`
  sex string `json:"性別"`
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
  if err != nil{
    //log.Fatal(err)
  }

  //クエリパラメータ
  params := request.URL.Query()
  params.Add("userId","Kuno") // 直接指定してみる
  request.URL.RawQuery = params.Encode()

  // ブラウザにjson表示
  //fmt.Fprintf(writer, )

  // Terminalにログ表示
  fmt.Println(params) // 'map[userId:[Kuno]]'
  fmt.Println(params["userId"]) // '[Kuno]'
  fmt.Println(request.URL.String())  // '?userId=Kuno'


  // MySQLからuserId=Kunoの情報取ってくる
  var query_sentence string
  query_sentence = "select * from user where user_id = '" + params["userId"] + "'"
  rows, err := db.Query(query_sentence)
  for rows.Next() {
    var user User
    err := rows.Scan(&user.user_id, &user.sex)

    //json.NewEncoder(w).Encode(user)
    // jsonエンコード
    outputJson, err := json.Marshal(&user)
    if err != nil {
      panic(err)
    }
    fmt.Fprint(w, string(outputJson))

    if err != nil {
      panic(err.Error())
      return
    }
    //fmt.Fprintf(w, "user_id : " + user.user_id)
    //fmt.Fprintf(w, ", sex : " + user.sex)
  }

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUser")
}
