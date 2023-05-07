.DEFAULT_GOAL := help

build-local: ## Build docker image to local development
	docker compose build --no-cache
.PHONY: build-local

up: ## Do docker compose up
	docker compose up -d
.PHONY: up

down: ## Do docker compose down
	docker compose down
.PHONY: down

logs: ## Tail docker compose logs
	docker compose logs -f
.PHONY: logs

ps: ## Check container status
	docker compose ps
.PHONY: ps

create-volume: ## Create docker volume for db
	docker volume create go-intermediate-practice-volume
.PHONY: create-volume

remove-volume: ## Remove docker volume for db
	docker volume rm go-intermediate-practice-volume
.PHONY: remove-volume

show-data: ## Show db data
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp -e "SELECT * FROM articles LEFT OUTER JOIN comments ON articles.article_id = comments.article_id"
.PHONY: show-data

seed:  ## Seed data
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp < /app/docker/mysql/init.d/1_createTable.sql
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp < /app/docker/mysql/init.d/2_createTable.sql
.PHONY: seed

load: ## Load server
	go run main.go
.PHONY: load

hot-reload: ## Hot-reload server
	air -c .air.toml
.PHONY: hot-reload

format-swagger: ## Format swagger annotation
	swag fmt
.PHONY: format-swagger

update-swagger: ## Update swagger.yaml
	swag init --outputTypes yaml
.PHONY: update-swagger

godoc: ## Load godoc
	godoc -http=:6060
.PHONY: godoc

test: ## Execute tests
	go test -v ./...
.PHONY: test

test-cover: ## Execute tests with coverage
	go test -v -cover ./...
.PHONY: test-cover

test-race: ## Execute tests (-race)
	go test -race -shuffle=on ./...
.PHONY: test-race

generate: ## Generate codes
	go generate ./...
.PHONY: generate

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
.PHONY: help
