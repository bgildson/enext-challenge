FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go run ./cmd/parser/main.go -log ./games.log -out ./games.json

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/api/main.go

FROM alpine:latest

COPY --from=builder /app/main /app/games.json /

EXPOSE 80

ENTRYPOINT [ "/main", "-games-json-path", "/games.json", "-port", "80" ]
