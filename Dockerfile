FROM golang:1.15
WORKDIR $GOPATH/src/github.com/eduardobobato/crud-go
COPY . .
RUN go get -d -v ./...;go install -v ./...;
EXPOSE 3333
ENTRYPOINT ["crud-go"]