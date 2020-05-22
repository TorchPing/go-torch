# syntax=docker/dockerfile:experimental
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.13-alpine as builder

ARG VERSION

ENV PORT=8080 \
         GIN_MODE=production
# Convert TARGETPLATFORM to GOARCH format
# https://github.com/tonistiigi/xx
COPY --from=tonistiigi/xx:golang / /

ARG TARGETPLATFORM

RUN apk --update --no-cache add \
    build-base \
    gcc \
    git \
  && rm -rf /tmp/* /var/cache/apk/*

ADD . /src

WORKDIR /src

ENV GO111MODULE=on

RUN  go env && env $(cat /tmp/.env | xargs) go build -ldflags "-w -s -X 'main.version=${VERSION}'"  -v -o torch cmd/main.go

FROM alpine:latest

WORKDIR /bin/

COPY --from=builder /src/torch ./torch

ENTRYPOINT ["/bin/torch"]