FROM golang:1.11.2-alpine3.8 AS build

WORKDIR /
COPY . /go/src/github.com/hisamura333/books-management
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/google/uuid \
  && go get github.com/gorilla/mux

CMD ["go", "run", "/go/src/github.com/hisamura333/books-management/main.go"]
