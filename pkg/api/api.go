// エンドポイントの定義とサーバ起動

package api

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
)

func Start() {
    http.HandleFunc("/", Handler)
    // 8080ポートでサーバ起動
    http.ListenAndServe(":8080", nil)
}

func Handler(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintf(writer, "Hello World!")
}
