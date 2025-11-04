# City Service

[🇬🇧 English version](README.md) | [🇷🇺 Русская версия](README_ru.md)

## Описание

Микросервис городов, написан на Golang (go 1.25.2, gin 1.11.0, gorm 1.31.0, postgres 16)

Сервис предоставляет API к библиотеке городов Российской Федерации.

Архитектура проекта основана на https://github.com/golang-standards/project-layout/blob/master/README_ru.md

## Развертывание

Развертывание проекта возможно двумя способами: 
1. локально (с установленным Golang на сервере)
2. с помощью утилиты Docker-compose.

### Локальный способ

Для работы необходим установленный Golang версии 1.25.2 и следующие пакеты:
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get github.com/joho/godotenv
```
Для запуска приложения используйте следующую команду:
```shell
 go run cmd/main.go
```

### способ с Docker-compose
Docker-контейнер необходимо запустить из корня проекта с помощью команды:
```shell
docker-compose up
```

Проект будет доступен по адресу: http://localhost:8080

## Makefile
makefile — это файл, содержащий набор инструкций, используемых утилитой make в инструментарии автоматизации сборки.

В проекте makefile имеет следующие команды:

```shell
init # сборка конфигов (на данный момент - перенос файлов из .env.example в .env)
up # Запуск контейнера с проектом
down # Остановка контейнера с проектом
```

## API

API предоставляет базовые CRUD для городов и регионов, а также метод синхронизации базы данных из файла в формате CVS
Описание API в формате OpenApi для Swagger доступно по следующей ссылке: [ссылка](https://github.com/PRodionovDev/CityService/blob/main/doc/openapi.yaml)

### Города
методы для городов:

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
методы для регионов:

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
синхронизация городов и регионов из формата CSV в базу данных:

```shell
curl -X POST http://localhost:8080/sync -H "Content-Type: application/json" -H "Authorization: test"
```

## Автотесты

В РАЗРАБОТКЕ