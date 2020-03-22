// エンドポイントの定義とサーバ起動
// エンドポイント：特定のリソースに対して与えられた固有の一意なURIのこと

package api

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"

  "github.com/YawarakasaApp/pkg/api/handler/user"
)

func Start() {
    http.HandleFunc("/", Handler)
    // 8080ポートでサーバ起動
    http.ListenAndServe(":8080", nil)
}

func Handler(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintf(writer, "Hello World!")
    // エンドポイントの設定
    // HandlerPkg　packageのGetUser関数呼び出し
    http.HandleFunc("/user/get", HandlerPkg.GetUser)
}
