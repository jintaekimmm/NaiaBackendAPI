# Naia Backend API

Naia Proejct Backend API

ElasticSearch를 조회하여 단어 순위등의 결과를 제공한다

- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [ElasticSearch Client](https://github.com/elastic/go-elasticsearch)
- [Wire](https://github.com/google/wire)
- [Swagger](https://github.com/swaggo/swag)

### Project Flow
```shell
Controllers -> Services -> Repositories -> Models 
```

### Set Environment
```shell
cp env_exampl env.sh

# changed ElasticSearch, index...
# and apply the modified environment variables
source env.sh
```
### Swagger
After modify swagger options

Swagger URL: http://IP:8000/swagger/index.html
```shell
go get -u github.com/swaggo/swag/cmd/swag
swag init -parseDependency=true
```

### Pre-Run(build)
Usage Wire(Dependency Injection)
```shell
$ wire
```

## Run
No build
```shell
go run main.go wire_gen.go
```

build
```shell
go build main.go wire_gen.go
./main
```