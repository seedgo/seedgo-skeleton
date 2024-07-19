FROM golang:1.20.4-alpine3.18 AS builder

ARG GOPROXY=https://goproxy.cn,direct
ARG GO111MODULE=on
ARG ALPINE_PKG_SRC="mirrors.aliyun.com"
ARG ALPINE_PKG_BASE="git"

ENV GOPROXY=$GOPROXY
ENV GO111MODULE=$GO111MODULE

WORKDIR /

COPY . .

RUN sed -i "s/dl-cdn.alpinelinux.org/$ALPINE_PKG_SRC/g" /etc/apk/repositories && \
    apk add --update --no-cache ${ALPINE_PKG_BASE}

RUN --mount=type=cache,target=/go/pkg/mod,id=gopkg \
--mount=type=cache,target=/root/.cache/go-build,id=gobuildcache \ 
go mod download \
&& go build -o app .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app /app/main

ENTRYPOINT ["/app/main"]
