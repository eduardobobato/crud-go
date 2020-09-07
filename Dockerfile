FROM golang
WORKDIR /app
COPY . .
RUN go get github.com/gorilla/mux;go get go.mongodb.org/mongo-driver/mongo;go get github.com/BurntSushi/toml;go get github.com/eduardobobato/crud-go/model;go get github.com/eduardobobato/crud-go/config;go get github.com/eduardobobato/crud-go/config/dao;go get github.com/eduardobobato/crud-go/service;go get github.com/eduardobobato/crud-go/router
RUN go build -o main .
EXPOSE 3333
ENTRYPOINT ["./main"]
