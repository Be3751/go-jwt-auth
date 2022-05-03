# Goの公式イメージ
FROM golang:1.17.8-alpine3.15

# docker run時に作業ディレクトリの作成とgitのインストールを実行
RUN mkdir /go/src/go-jwt-auth
RUN apk update && apk add git

# ワーキングディレクトリの指定
WORKDIR /go/src/go-jwt-auth

# ホストのファイル・ディレクトリをコンテナイメージにコピー
ADD . /go/src/go-jwt-auth

RUN go mod download

# ホットリロード用にairを導入
RUN go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]