# Naia Backend API

Naia Proejct Backend APIs

- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [ElasticSearch Client](https://github.com/elastic/go-elasticsearch)
- [Redis Client](https://github.com/go-redis/redis)
- [Swagger](https://github.com/swaggo/swag)


## Set Environment
```shell
$ cp env_example env.sh

# changed ElasticSearch, index...
# and apply the modified environment variables
$ source env.sh
```
## Swagger
After modify swagger options

Swagger URL: http://IP:8000/swagger/index.html
```shell
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag init -parseDependency=true
```

## Run
No build
```shell
$ go run main.go
```

build
```shell
$ go build main.go
$ ./main
```

## Docker
Image Build
```shell
$ docker build -t [CONTAINER_REPOSITORY]:[TAG] \
 --build-arg ELS_HOST[ELS_HOST] \
 --build-arg ELS_INDEX=[ELS_INDEX] \
 --build-arg ELS_USER=[ELS_USER] \
 --build-arg ELS_PASSWORD=[ELS_PASSWORD] \    
 --build-arg REDIS_HOST=[REDIS_HOST] \
 --build-arg REDIS_PORT=[REDIS_PORT] \
 --build-arg REDIS_DB=[REDIS_DB] \
 --build-arg REDIS_PASSWORD=[REDIS_PASSWORD] \
 --build-arg REDIS_KEY=[REDIS_KEY] \
 --build-arg RELATED_API=[RELATED_API] \
 .
```

Container Run
```shell
$ docker run -d --name [CONTAINER_NAME] -p 8000:8000 --restart=always [CONTAINER_REPOSITORY]:[CONTAINER_VERSION]
```
