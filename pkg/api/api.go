// エンドポイントの定義とサーバ起動
// エンドポイント：特定のリソースに対して与えられた固有の一意なURIのこと

package api

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "net/http"

  "github.com/YawarakasaApp/pkg/api/handler/user"
  "github.com/YawarakasaApp/pkg/api/middleware"
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
    //http.HandleFunc("/user/get", HandlerPkg.GetUser)
    http.HandleFunc("/user/create", HandlerPkg.HandleUserCreate)

    // こうなってほしい
    //http.HandleFunc("/user/get",　get(middleware.Authenticate(handler.HandleUserGet())))

    //http.HandleFunc("/user/get",　get(middleware.Authenticate(HandlerPkg.GetUser))
    http.HandleFunc("/user/get", middleware.Authenticate(HandlerPkg.HandleUserGet))
    //http.HandleFunc("/user/update",post(middleware.Authenticate(handler.HandleUserUpdate())))
}

/*
func get(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
      fmt.Println("[START] get")
      next.ServeHTTP(w, r)
  }
}
*/
