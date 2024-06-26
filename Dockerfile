FROM golang:1.22.4-bullseye

RUN mkdir -p /app
WORKDIR /app
COPY . .

# airのインストール
RUN go install github.com/air-verse/air@latest

# 必要なモジュールのダウンロード
RUN go mod download

# airの設定ファイルをコピー
COPY .air.toml /app/.air.toml

# ポートの公開
EXPOSE 8080

# airを使ってアプリケーションを起動
CMD ["air"]