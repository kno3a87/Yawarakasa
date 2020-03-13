package main

import (
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, Kuno")
  fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
  // API起動
  handleRequests()
}
