FROM golang:1.20.4-alpine3.18 AS builder

WORKDIR /

COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w  GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
RUN go mod download

RUN go build -o app .

FROM alpine:3.18.0

WORKDIR /app

COPY --from=builder /app /app/main

ENTRYPOINT ["/app/main"]
