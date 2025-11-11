init:
	$(shell sed 's/%APP_ID%/service/g' .env.example > .env)
up:
	docker-compose up -d --build --remove-orphans

down:
	docker-compose down --remove-orphans
tests:
	docker exec -it city-service-container go test ./test/unit