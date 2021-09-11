#!/bin/bash -x

# goアプリケーションのビルド
go build -o first_webapp_morita main.go

# goアプリケーションの実行
./first_webapp_morita
