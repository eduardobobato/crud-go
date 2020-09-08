# crud-go

### Instalando dependencias
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get github.com/BurntSushi/toml
go get github.com/eduardobobato/crud-go/model
go get github.com/eduardobobato/crud-go/config
go get github.com/eduardobobato/crud-go/dao
go get github.com/eduardobobato/crud-go/service
go get github.com/eduardobobato/crud-go/controller

### Rodando App
go run .\main.go

### Generate Swagger
* go get -u github.com/go-swagger/go-swagger/cmd/swagger
* swagger generate spec -o ./swagger.yaml --scan-models

### Rondando no Docker
docker build -f "Dockerfile" -t crud-go:1.0.0 .
docker run -d -p 3333:3333 crud-go:1.0.0