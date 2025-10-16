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

## ToDo
1. расширить структуру городов и регионов (координаты и тд)
2. авторизация
3. заполнение дампа