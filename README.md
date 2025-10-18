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
curl -X GET http://localhost:8080/cities -H "Content-Type: application/json" -H "Authorization: test"
```
```shell
curl -X GET http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```
```shell
curl -X POST http://localhost:8080/city -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```
```shell
curl -X PUT http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```
```shell
curl -X DELETE http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```
### Regions
```shell
curl -X GET http://localhost:8080/regions -H "Content-Type: application/json" -H "Authorization: test"
```
```shell
curl -X GET http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```
```shell
curl -X POST http://localhost:8080/region -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```
```shell
curl -X PUT http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```
```shell
curl -X DELETE http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

## ToDo
1. заполнение дампа
2. расписать проект в readme