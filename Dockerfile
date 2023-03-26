FROM golang:1.20.2

# dockerもGoの作業ディレクトリの概念がわからん
WORKDIR /go/src/todo_app_backend

COPY . .