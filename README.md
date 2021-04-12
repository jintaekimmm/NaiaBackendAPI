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

## Docker
Docker Build
```shell
 docker build -t [CONTAINER_NAME]:[TAG] \
 --build-arg ELS_HOST=http://IP:9200 \
 --build-arg ELS_INDEX=[INDEX] \
 --build-arg ELS_USER=[USERNAME] \
 --build-arg ELS_PASSWORD=[PASSWORD] \    
 --build-arg REDIS_HOST="${REDIS_HOST}" \
 --build-arg REDIS_PORT="${REDIS_PORT}" \
 --build-arg REDIS_DB="${REDIS_DB}" \
 --build-arg REDIS_PASSWORD="${REDIS_PASSWORD}" \
 --build-arg REDIS_KEY="${REDIS_KEY}" .
```

Docker run
```shell
docker run -d --name [CONTAINER_NAME] -p 8000:8000 --restart=always [CONTAINER_NAME]:1.0
```