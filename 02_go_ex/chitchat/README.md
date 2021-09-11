# ChitChat
こちらのソースコードをGoの勉強のために使わせていただいております。
https://github.com/mushahiroyuki/gowebprog/tree/master/ch02/chitchat

# 実行手順
```sh
$ ＜Postgresを起動＞
$ createdb chitchat
$ psql -f data/setup.sql -d chitchat
$ go build
$ ./chitchat
```
