# echoフレームワークとクリーンアーキテクチャのサンプルアプリ

こちらの資料を元に実装させていただきました。
https://qiita.com/mIchino/items/b885de3396e3f77d8b37

## 実行手順
### DBの準備
```
$ cd 02_go_ex
$ pwd
  -- <ワークスペース>/one-more-advice-prog/02_go_ex
$ docker-compose up -d one_more_advice_db_prog
or
$ docker-compose up one_more_advice_db_prog
```

DBに接続
```
- ホスト
  - 127.0.0.1
- ユーザー
  - root
- パスワード
  - one_more_advice_prog
- DB名
  - one_more_advice_prog
- ポート
  - 4307
```

### webアプリの実行
```sh
$ cd 02_go_ex/echo_prog/02_echo_sample_api/src
$ pwd
  -- <ワークスペース>/one-more-advice-prog/02_go_ex/echo_prog/02_echo_sample_api/src

# 初回のみ
$ go mod init echo_sample_api
$ go get github.com/labstack/echo
$ go get github.com/labstack/echo/v4
$ go get gorm.io/driver/mysql
$ go get gorm.io/gorm
# webアプリの実行
$ go run server.go
```

## 動作確認
### POST: /users
```sh
$ curl --location --request POST 'localhost:1323/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"J.Y Park"
}'


$ curl --location --request POST 'localhost:1323/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Eron Mask"
}'
```

### GET: /users
```sh
$ curl --location --request GET 'localhost:1323/users'
```

### DELETE: /users/:id
```sh
$ curl --location --request DELETE 'localhost:1323/users/1'
```
