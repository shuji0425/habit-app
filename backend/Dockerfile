# ベースイメージ
FROM golang:1.23

# 作業ディレクトリ
WORKDIR /app/backend

# モジュールのキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースのコピー
COPY . .

# airをインストール
RUN go install github.com/air-verse/air@latest

# 初回のairが始まるまで少し待つので先にbuild
# ビルド
RUN go build -o main .

CMD ["air"]