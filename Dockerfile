FROM golang:1.10-alpine AS build

RUN apk add --update curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/alcmoraes/go-data-integration-challenge
WORKDIR /go/src/github.com/alcmoraes/go-data-integration-challenge

COPY . .

RUN dep ensure

RUN go build -o /app/api

CMD ["/app/api"]
