init:
	$(shell sed 's/%APP_ID%/service/g' .env.example > .env)
up:
	docker-compose up -d --build --remove-orphans

down:
	docker-compose down --remove-orphans
tests: