FROM golang:1.22.4-bullseye

RUN mkdir -p /app
WORKDIR /app

# go.modとgo.sumをコピー
COPY go.mod go.sum ./

# 必要なモジュールのダウンロード
RUN go mod download

# 残りのファイルをコピー
COPY . .

# airのインストール
RUN go install github.com/air-verse/air@latest

# airの設定ファイルをコピー
COPY dev/.air.toml /app/.air.toml

# ポートの公開
EXPOSE 8080

# airを使ってアプリケーションを起動
CMD ["air"]