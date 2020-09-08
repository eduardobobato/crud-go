FROM golang:1.15
WORKDIR $GOPATH/src/github.com/eduardobobato/crud-go
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .
EXPOSE 3333
ENTRYPOINT ["go-sample-app"]