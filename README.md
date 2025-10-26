# City Service

## Description

Cities microservice with a Go stack (go 1.25.2, gin 1.11.0, gorm 1.31.0, sqlite)

The service provides an API to the library of cities of the Russian Federation.

## Installation

Clone the repository into the new project directory

```shell
mkdir city_service
cd city_service
git clone git@github.com:PRodionovDev/CityService.git
```

## Deployment

Deployment is possible in two ways: locally (with Golang installed on the server) or using a Docker app

### Local

You need to install Golang version 1.25.2 and the following packages:
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/joho/godotenv
```
To run use:
```shell
 go run cmd/main.go
```

### Docker
The Docker container must be launched from the project root using the command:
```shell
docker-compose up
```

## Url:
The project is available at the URL: http://localhost:8080

## API

The API description in OpenApi format is available at the following link: [click here](https://github.com/PRodionovDev/CityService/blob/main/Doc/openapi.yaml)

### City
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
Basic CRUD for cities:

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