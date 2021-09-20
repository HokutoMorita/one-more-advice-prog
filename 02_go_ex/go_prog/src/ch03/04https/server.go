package main

import (
    "net/http"
)

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
        Handler: nil,
    }
    // cert.pemがSSL証明書
    // key.pemがサーバ用の秘密鍵
    server.ListenAndServeTLS("cert.pem", "key.pem")
}
