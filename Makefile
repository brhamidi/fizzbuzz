APP=fizzbuzz
D_PATH=Dockerfile

.PHONY: up dev down

## up the local stack
up:
	@D_PATH=$(D_PATH) docker-compose up --remove-orphans --build -d
	@docker-compose logs -f ${APP}

##	up local stack in development mode
##	a filewatcher is present for auto-reload the app
dev: D_PATH=Dockerfile.dev
dev: up

# down the local stack
down:
	@docker-compose down
