FROM golang

ADD . /go/src/crud-go

RUN go get github.com/gorilla/mux

RUN go install crud-go

ENTRYPOINT /go/bin/crud-go

EXPOSE 3333