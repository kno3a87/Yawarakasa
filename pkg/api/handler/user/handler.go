// 各エンドポイントの処理を定義


package HandlerPkg

import (
  //"encoding/json"
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

/*
  // queryにクエリパラメータ代入
  var query string
  query = "select * from user where user_id = '" + r.URL.RawQuery + "'"
  rows, err := db.Query(query)

  if err != nil {
    panic(err.Error())
    return
  }
  defer rows.Close()

  for rows.Next() {
    var user User
    err := rows.Scan(&user.user_id, &user.sex)

    json.NewEncoder(w).Encode(user)

    if err != nil {
      panic(err.Error())
      return
    }
    //fmt.Fprintf(w, "user_id : " + user.user_id)
    //fmt.Fprintf(w, ", sex : " + user.sex)
  }

  err = rows.Err()
  if err != nil {
    panic(err.Error())
    return
  }*/

  //url := "https://jsonplaceholder.typicode.com/todos"
  url := "" // ここなに？？？？？

	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		//log.Fatal(err)
	}

	//クエリパラメータ
	params := request.URL.Query()
    params.Add("userId","1")
    request.URL.RawQuery = params.Encode()

	fmt.Println(request.URL.String()) //https://jsonplaceholder.typicode.com/todos?userId=1

  // Terminalにログ表示
  fmt.Println("Endpoint Hit: drawUser")
}
