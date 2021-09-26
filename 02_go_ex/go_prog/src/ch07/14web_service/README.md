# Webサービス
```
$ cd 02_go_ex/go_prog/src/ch07/14web_service
```
## 初回のみ
```
$ go mod init api_server
$ go get github.com/go-sql-driver/mysql
```
## DBの準備
```
$ cd 02_go_ex
$ docker-compose up one_more_advice_db_prog
```

02_go_ex/go_prog/src/ch07/14web_service/setup.sql を実行してテーブルを準備する

## APIサーバの起動
```
$ go build -o api_server server.go
$ ./api_server
```
## データを登録
```
$ curl -i -X POST -H "Content-Type: application/json"  -d '{"content":"My first post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/
```
## データを更新
```
$ curl -i -X PUT -H "Content-Type: application/json"  -d '{"content":"Updated post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/1
```
## データの取得
```
$ curl -i -X GET http://127.0.0.1:8080/post/1
```
## データの削除
```
$ curl -i -X DELETE http://127.0.0.1:8080/post/1
```
