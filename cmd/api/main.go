// pkg/api/api.goのサーバ起動処理を呼び出すだけ

package main

// 別ファイル呼び出し（相対パスじゃだめ）
import "github.com/YawarakasaApp/pkg/api"

func main() {
  // api packageのStart関数呼び出し
  api.Start()
}
