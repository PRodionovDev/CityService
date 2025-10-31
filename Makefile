#golang app
init_go: update-configs \
  run

run:
	go run cmd/main.go

stop:
	kill -9 $$(lsof -t -i:8080)

tests:
	go run cmd/main.go --stand=test &
	go clean -testcache
	go test ./test
	kill -9 $$(lsof -t -i:8080)
	rm -rf pkg/database/dump/cities_test.db

#docker app
init: update-configs \
  up

update-configs:
    $(shell sed 's/%APP_ID%/service/g' .env.example > .env)

up:
	docker-compose up -d --build --remove-orphans

down:
	docker-compose down --remove-orphans
