FROM golang:1.18-alpine as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
apk add --no-cache $PACKAGES && make build

FROM alpine:3.10

COPY --from=builder /go/src/openapi /usr/local/bin
EXPOSE 30000
CMD openapi
