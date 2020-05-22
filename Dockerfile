# syntax=docker/dockerfile:experimental
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.13-alpine as builder

# Convert TARGETPLATFORM to GOARCH format
# https://github.com/tonistiigi/xx
COPY --from=tonistiigi/xx:golang / /

ARG TARGETPLATFORM

RUN apk add --no-cache musl-dev git gcc

ADD . /src

WORKDIR /src

ENV GO111MODULE=on

RUN cd cmd/torch && go env && go build -v

FROM alpine:latest

WORKDIR /bin/

COPY --from=builder /src/cmd/torch/torch .

ENTRYPOINT ["/bin/torch"]