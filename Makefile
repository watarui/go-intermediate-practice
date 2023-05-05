.DEFAULT_GOAL := help

.PHONY: build-local
build-local: ## Build docker image to local development
	docker compose build --no-cache

.PHONY: up
up: ## Do docker compose up
	docker compose up -d

.PHONY: down
down: ## Do docker compose down
	docker compose down

.PHONY: logs
logs: ## Tail docker compose logs
	docker compose logs -f

.PHONY: ps
ps: ## Check container status
	docker compose ps

.PHONY: create-volume
create-volume: ## Create docker volume for db
	docker volume create go-intermediate-practice-volume

.PHONY: remove-volume
remove-volume: ## Remove docker volume for db
	docker volume rm go-intermediate-practice-volume

.PHONY: test
test: ## Execute tests
	go test -race -shuffle=on ./...

.PHONY: show-data
show-data: ## Show db data
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp -e "SELECT * FROM articles LEFT OUTER JOIN comments ON articles.article_id = comments.article_id"

.PHONY: seed
seed:  ## Seed data
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp < /app/docker/mysql/init.d/1_createTable.sql
	mysql -u$(DB_USER) -h$(DB_HOST) -P3306 -D$(DB_NAME) -p$(DB_PASSWORD) --protocol=tcp < /app/docker/mysql/init.d/2_createTable.sql

.PHONY: generate
generate: ## Generate codes
	go generate ./...

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
