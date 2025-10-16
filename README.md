# City Service

## Назначение

Микросервис городов со стеком на Go (go 1.25.2, gin 1.11.0, gorm 1.31.0, sqlite)

## Развертывание

1. Клонируем репозиторий в директорию нового проекта

```shell
mkdir city_service
cd city_service
git clone git@github.com:PRodionovDev/CityService.git
```


## Локальный запуск

Запускаем инициализацию

```shell
go run main.go
```

## Docker

```shell
docker-compose up
```

Проект доступен по url:
http://localhost:8080

## API

Получение всех городов
```shell
curl -X GET http://localhost:8080/cities -H "Content-Type: application/json"
```
Получение города по ID
```shell
curl -X GET http://localhost:8080/city/{id} -H "Content-Type: application/json"
```
Создание нового города
```shell
curl -X POST http://localhost:8080/city -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
Редактирование города по ID
```shell
curl -X PUT http://localhost:8080/city/{id} -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
Удаление города по ID
```shell
curl -X DELETE http://localhost:8080/city/{id} -H "Content-Type: application/json"
```
Получение всех регионов
```shell
curl -X GET http://localhost:8080/regions -H "Content-Type: application/json"
```
Получение региона по ID
```shell
curl -X GET http://localhost:8080/region/{id} -H "Content-Type: application/json"
```
Создание нового региона
```shell
curl -X POST http://localhost:8080/region -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
Редактирование региона по ID
```shell
curl -X PUT http://localhost:8080/region/{id} -H "Content-Type: application/json" -d '{"name":"Наименование","slug":"Slug"}'
```
Удаление региона по ID
```shell
curl -X DELETE http://localhost:8080/region/{id} -H "Content-Type: application/json"
```

## ToDo
1. расширить структуру городов и регионов (координаты и тд)
2. добавить сваггер
3. русификация полная (или же наоборот)
4. авторизация
5. заполнение дампа