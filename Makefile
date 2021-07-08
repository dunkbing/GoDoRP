cnf ?= .env
include $(cnf)

# DOCKER TASKS
# Build the container
build: ## Build the container
	docker build -t $(APP_NAME) .

build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

run: ## Run container on port configured in `config.env`
	docker run -i -t --rm --env-file=./.env -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

clean:
	go clean
	sudo rm -rf ./db/pgdata