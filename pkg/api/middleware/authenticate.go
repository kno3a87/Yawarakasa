
package middleware

import (
  "context"
  "fmt"
  "net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("[START] middleware")

        url := ""
        request, err := http.NewRequest("GET", url, nil)
        // リクエストヘッダに代入
        request.Header.Set("userId", "Kuno")
        if err != nil{
          //log.Fatal(err)
        }
        // リクエストヘッダから取得してctxに保存
        // 空のコンテキスト作成
        Ctx := context.Background()
        // コンテキストに代入
        Ctx = context.WithValue(Ctx, "userId", request.Header.Get("userId"))
        // Terminalにログ表示
        fmt.Println(request.Header.Get("userId"))
        fmt.Println(Ctx.Value("userId"))

        // HandlerPkg.GetUserへ
        next.ServeHTTP(w, r)

        fmt.Println("[END] middleware")
    }
}
