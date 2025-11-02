# City Service

[🇬🇧 English version](README.md) | [🇷🇺 Русская версия](README_ru.md)

## Description

Cities microservice with a Go stack (go 1.25.2, gin 1.11.0, gorm 1.31.0, sqlite)

The service provides an API to the library of cities of the Russian Federation.

The project architecture is based on https://github.com/golang-standards/project-layout

## Deployment

There are two ways to deploy a project:
1. Locally (with Golang installed on the server)
2. Using the Docker-compose utility.

### Local

You need to install Golang version 1.25.2 and the following packages:
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/joho/godotenv
```
To run the application use the following command:
```shell
 go run cmd/main.go
```

### Docker
The Docker container must be launched from the project root using the command:
```shell
docker-compose up
```

The project is available at the URL: http://localhost:8080

## Makefile
A makefile is a file containing a set of instructions used by the make utility in the build automation toolchain.

In a project, the makefile contains the following commands:

1. for local use:
```shell
init_go # project initialization (building configs and launching)
run # running app
stop # Stopping the project by killing the process that is using port 8080
tests # Running integration tests (launching the application with the test environment, resetting the cache, running tests, shutting down the application, and deleting the test database)
```

2. for use with Docker:
```shell
init # project initialization (building configs and launching)
up # up Docker container
down # stop Docker container
```

3. And also the general command:
```shell
update-configs # Building configs (currently moving files from .env.example to .env)
```

## API

The API description in OpenApi format is available at the following link: [click here](https://github.com/PRodionovDev/CityService/blob/main/doc/openapi.yaml)

The API provides basic CRUD for cities and regions, as well as a method for synchronizing the database from a CVS file.
The API description in OpenApi format for Swagger is available at the following link: [ссылка](https://github.com/PRodionovDev/CityService/blob/main/doc/openapi.yaml)

### Cities
Basic CRUD for cities:

1. get all cities
```shell
curl -X GET http://localhost:8080/cities -H "Content-Type: application/json" -H "Authorization: test"
```

2. get city by id
```shell
curl -X GET http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

3. create new city
```shell
curl -X POST http://localhost:8080/city -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```

4. update city by id
```shell
curl -X PUT http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```

5. delete city by id
```shell
curl -X DELETE http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```
### Regions
Basic CRUD for regions:

1. get all regions
```shell
curl -X GET http://localhost:8080/regions -H "Content-Type: application/json" -H "Authorization: test"
```

2. get region by id
```shell
curl -X GET http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

3. create new region
```shell
curl -X POST http://localhost:8080/region -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```

4. update region by id
```shell
curl -X PUT http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```

5. delete region by id
```shell
curl -X DELETE http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

### Sync
Basic synchronization of cities and regions from CSV format to a database:

```shell
curl -X POST http://localhost:8080/sync -H "Content-Type: application/json" -H "Authorization: test"
```

## Test

Automated tests are implemented as integration tests that check API requests with an empty test database.

To run automated tests, use the following algorithm:

```shell
go run cmd/main.go --stand=test #launch the application with a test environment
```
```shell
go clean -testcache #clean cache
go test ./test #run tests
```
```shell
rm -rf pkg/database/dump/cities_test.db #delete the test database dump
```

### Roadmap

1. PostgreSQL Docker
2. Connect to Docker PostgreSQL
3. Autotest with Docker