FROM golang:1.16.0-alpine

ENV GOPATH /go
ENV GO111MODULE=on

RUN mkdir -p /go/src/app
COPY ./go.mod /go/src/app
COPY ./go.sum /go/src/app
WORKDIR /go/src/app

RUN go mod download
#RUN go mod tidy

RUN set -eux && \
    apk update && \
    go get -u github.com/cosmtrek/air
    #go get -u github.com/cosmtrek/air \
    #go build -o /go/bin/air github.com/cosmtrek/air


EXPOSE 8080
COPY ./ /go/src/app

ENTRYPOINT air -c ./graph/.air.toml
