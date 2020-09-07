# crud-go

### Instalando dependencias
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get github.com/BurntSushi/toml
go get github.com/eduardobobato/crud-go/model
go get github.com/eduardobobato/crud-go/config
go get github.com/eduardobobato/crud-go/config/dao
go get github.com/eduardobobato/crud-go/service
go get github.com/eduardobobato/crud-go/router

### Rodando App
go run .\main.go

### Generate Swagger
* swagger generate spec -o ./swagger.yaml --scan-models