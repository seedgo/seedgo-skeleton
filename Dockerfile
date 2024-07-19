FROM golang:1.20.4-alpine3.18 AS builder

ARG GOPROXY=https://goproxy.cn,direct
ARG GO111MODULE=on

ENV GOPROXY=$GOPROXY
ENV GO111MODULE=$GO111MODULE

WORKDIR /

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod,id=gopkg \
--mount=type=cache,target=/root/.cache/go-build,id=gobuildcache \ 
go mod download \
&& go build -o app .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app /app/main

ENTRYPOINT ["/app/main"]
