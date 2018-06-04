FROM golang:1.10-alpine AS build

RUN apk add --update curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/alcmoraes/go-data-integration-challenge
WORKDIR /go/src/github.com/alcmoraes/go-data-integration-challenge

COPY . .

RUN curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
RUN chmod +x ./cc-test-reporter

RUN go get github.com/axw/gocov/gocov

RUN dep ensure

RUN mkdir /app

RUN go build -o /app/api

COPY config.yml.dev config.yml

CMD ["/app/api"]
