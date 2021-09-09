# Go言語の学習

## Docker環境構築コマンド
```sh
# Dockerビルド
$ docker-compose build go_prog
$ docker-compose run go_prog sh
```

## Goの実行コマンド
```sh
# コンパイル & 実行
$ go run main.go

# mainパッケージのコードを複数実行
$ go run main.go app.go
$ go run *.go
```


### プログラムのビルド
```sh
# コンパイルして実行ファイルを生成
## -oオプションを使用して出力する実行ファイルのファイル名を指定できる
### helloアプリケーションの場合
$ pwd
  - /go/app/src/hello
$ go build -o hello main.go

### zooアプリケーションの場合
$ pwd
  - /go/app/src/zoo
$ go build -o zoo main.go
```
#### 複数のmainパッケージのコードをビルド
```sh
$ go build
$ go build *.go
$ go build -o zoo *.go
```


### 実行ファイルからプログラムを実行
```sh
# helloアプリケーションの場合
$ pwd
  - /go/app/src/hello
$ ./hello

# zooアプリケーションの場合
$ pwd
  - /go/app/src/zoo
$ ./zoo
```

### Moduleパッケージのテストの実行 
- ファイル名に、_test.goというサフィックスがあることでテストファイルと認識される
- 関数名に、Testというプレフィックスがあることでテスト関数と認識される
```sh
$ go test ./animals

# 個々のテストの実行詳細を確認する場合は-vオプションを付ける
$ go test -v ./animals
```

### mainパッケージのテストの実行
- ファイル名に、_test.goというサフィックスがあることでテストファイルと認識される
- 関数名に、Testというプレフィックスがあることでテスト関数と認識される
```sh
$ go test
$ go test -v
```

## Goのメタ情報確認
### Goの実行環境のバージョン確認
```sh
$ go version
```

### Goの実行環境の確認
```sh
$ go env
```

### GoのModulesの初期化
```
$ go mod init zoo
```

## 参考資料
  - 【Go言語】作成したpackageをimportして使用。「cannot find module for path」エラー時の解決
    - https://qiita.com/namari/items/a037e2167eeec3a20ea7
  - Go言語の依存モジュール管理ツール Modules の使い方
    - https://qiita.com/uchiko/items/64fb3020dd64cf211d4e
  - Go Modulesも触れてみるGo入門
    - https://tech.opst.co.jp/2019/07/09/go-modules%E3%82%82%E8%A7%A6%E3%82%8C%E3%81%A6%E3%81%BF%E3%82%8Bgo%E5%85%A5%E9%96%80/
## 用語
  - 
