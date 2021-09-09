# goバージョン
FROM golang:1.16.3-alpine

# アップデートとgitのインストール
RUN apk add --update && apk add git

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /go/app

# ワーキングディレクトリの設定
WORKDIR /go/app

