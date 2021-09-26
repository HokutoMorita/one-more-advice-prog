# Webサービス
$ cd 02_go_ex/go_prog/src/ch07/14web_service
## 初回のみ
$ go mod init api_server
## DBの準備
$ cd 02_go_ex
$ docker-compose up one_more_advice_db_prog
## APIサーバの起動
$ go build -o api_server server.go
$ ./api_server
