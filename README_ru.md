# City Service
[🇬🇧 English version](README.md) | [🇷🇺 Русская версия](README_ru.md)
## Описание

Микросервис город на Golang (go 1.25.2, gin 1.11.0, gorm 1.31.0, sqlite)
Сервис предоставляет API к библиотеке городов Российской Федерации.
Архитектура проекта основана на https://github.com/golang-standards/project-layout/blob/master/README_ru.md

## Установка

Клонируйте репозиторий в новый каталог проекта.

```shell
mkdir city_service
cd city_service
git clone git@github.com:PRodionovDev/CityService.git
```

## Развертывание

Развертывание возможно двумя способами: локально (с установленным Golang на сервере) или с помощью приложения Docker.

### Локальное

Вам необходимо установить Golang версии 1.25.2 и следующие пакеты:
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/joho/godotenv
```
Для запуска используйте:
```shell
 go run cmd/main.go
```

### Docker
Docker-контейнер необходимо запустить из корня проекта с помощью команды:
```shell
docker-compose up
```

## Url:
Проект доступен по URL: http://localhost:8080

## API

Описание API в формате OpenApi доступно по следующей ссылке: [ссылка](https://github.com/PRodionovDev/CityService/blob/main/doc/openapi.yaml)

### Города
Базовый CRUD для городов:

1. Получить все города
```shell
curl -X GET http://localhost:8080/cities -H "Content-Type: application/json" -H "Authorization: test"
```

2. Получить город по id
```shell
curl -X GET http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

3. Создать новый город
```shell
curl -X POST http://localhost:8080/city -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```

4. Обновить город по id
```shell
curl -X PUT http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Наименование","slug":"Slug","region_id":1,"is_capital":true,"type":"Город","latitude":55.751244,"longitude":37.618423,"time_zone":"Europe/Moscow","population":13274285}'
```

5. Удалить город по id
```shell
curl -X DELETE http://localhost:8080/city/{id} -H "Content-Type: application/json" -H "Authorization: test"
```
### Регионы
Базовый CRUD для регионов:

1. Получить все регионы
```shell
curl -X GET http://localhost:8080/regions -H "Content-Type: application/json" -H "Authorization: test"
```

2. Получить регион по id
```shell
curl -X GET http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

3. Добавить новый регион
```shell
curl -X POST http://localhost:8080/region -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```

4. Обновить регион по id
```shell
curl -X PUT http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test" -d '{"name":"Москва","slug":"Moscow","number":99}'
```

5. Удалить регион по id
```shell
curl -X DELETE http://localhost:8080/region/{id} -H "Content-Type: application/json" -H "Authorization: test"
```

### Синхронизация
Базовая синхронизация городов и регионов из формата CSV в базу данных:

```shell
curl -X POST http://localhost:8080/sync -H "Content-Type: application/json" -H "Authorization: test"
```

## Автотесты

```shell
go clean -testcache
go test ./test
```
