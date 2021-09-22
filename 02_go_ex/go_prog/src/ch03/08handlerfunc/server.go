package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello! morita")
}

func world(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "World! morita")
}

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    // 関数をハンドラに変換して、DefaultServeMuxに登録する関数 => http.HandleFunc()
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/world", world)

    server.ListenAndServe()
}
