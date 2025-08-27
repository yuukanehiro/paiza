FROM golang:1.24-alpine

WORKDIR /app

# 必要ツール
RUN apk add --no-cache bash tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && echo "Asia/Tokyo" > /etc/timezone

# 依存解決のため先にgo.modだけCOPY→go mod download（ビルドキャッシュ効かせる）
COPY go.mod ./
RUN go mod download

# 残りのソース
COPY . .