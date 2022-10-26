FROM golang:1.19-alpine3.16 AS builder

WORKDIR /usr/app

COPY . .

RUN apk add --update gcc g++
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o server -ldflags="-w -s" src/main.go

EXPOSE 8080

ENTRYPOINT [ "/usr/app/server" ]

