APP=fizzbuzz

.PHONY: up

up:
	@docker-compose up --remove-orphans --build -d
	@docker-compose logs -f

down:
	@docker-compose down
