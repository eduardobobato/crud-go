FROM golang
ADD . /go/src/crud-go
RUN go get github.com/gorilla/mux;go get go.mongodb.org/mongo-driver/mongo;go get github.com/BurntSushi/toml;go get github.com/eduardobobato/crud-go/model;go get github.com/eduardobobato/crud-go/config;go get github.com/eduardobobato/crud-go/config/dao;go get github.com/eduardobobato/crud-go/service;go get github.com/eduardobobato/crud-go/router
run go build main.go
ENTRYPOINT /go/src/crud-go/crud-go

EXPOSE 3333