# City Service

## Description

Cities microservice with a Go stack (go 1.25.2, gin 1.11.0, gorm 1.31.0, sqlite)

## Deployment

Clone the repository into the new project directory

```shell
mkdir city_service
cd city_service
git clone git@github.com:PRodionovDev/CityService.git
```

### Local

```shell
go run main.go
```

### Docker

```shell
docker-compose up
```

### Url:
http://localhost:8080

## API

[OpenApi Yaml](https://github.com/PRodionovDev/CityService/blob/main/Doc/openapi.yaml)

### City
```shell
curl -X GET http://localhost:8080/cities -H "Content-Type: application/json"
```
```shell
curl -X GET http://localhost:8080/city/{id} -H "Content-Type: application/json"
```
```shell
curl -X POST http://localhost:8080/city -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
```shell
curl -X PUT http://localhost:8080/city/{id} -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
```shell
curl -X DELETE http://localhost:8080/city/{id} -H "Content-Type: application/json"
```
### Regions
```shell
curl -X GET http://localhost:8080/regions -H "Content-Type: application/json"
```
```shell
curl -X GET http://localhost:8080/region/{id} -H "Content-Type: application/json"
```
```shell
curl -X POST http://localhost:8080/region -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
```shell
curl -X PUT http://localhost:8080/region/{id} -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
```shell
curl -X DELETE http://localhost:8080/region/{id} -H "Content-Type: application/json"
```

## ToDo
1. расширить структуру городов и регионов (координаты и тд)
2. авторизация
3. заполнение дампа