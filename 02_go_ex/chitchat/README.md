
# ChitChat
こちらのソースコードをGoの勉強のために使わせていただいております。
https://github.com/mushahiroyuki/gowebprog/tree/master/ch02/chitchat

# 実行手順
## 1. DBの準備
下記の手順でコマンドを実行
```sh
$ cd <ワークスペース>/one-more-advice-prog/02_go_ex
$ pwd
  - <ワークスペース>/one-more-advice-prog/02_go_ex
$ docker-compose up -d one_more_advice_db_prog
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

DB内で下記のクエリを実行
  - 02_go_ex/chitchat/data/setup.sql

## 2. サンプルアプリの立ち上げ
下記の手順でコマンドを実行
```sh
$ cd <ワークスペース>/one-more-advice-prog/02_go_ex/chitchat
$ go build
$ ./chitchat
```

## 3. ブラウザで確認
  - http://localhost:8080


![スクリーンショット 2021-09-20 10 34 54](https://user-images.githubusercontent.com/38268537/133950389-936add6c-7de8-45ed-a714-373cd572ba75.png)
