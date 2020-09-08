# crud-go

### Instalando dependencias
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/mongo
go get github.com/BurntSushi/toml
go get github.com/stretchr/testify

### Rodando App
go run .\main.go

### Rodando os testes
go test ./service/... ./controller/...

### Generate Swagger
* go get -u github.com/go-swagger/go-swagger/cmd/swagger
* swagger generate spec -o ./swagger.yaml --scan-models

### Rondando no Docker
docker build -f "Dockerfile" -t crud-go:1.0.0 .
docker run -d -p 3333:3333 crud-go:1.0.0

#### Com Swarm
docker-compose up --build