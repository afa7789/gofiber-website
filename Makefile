#!/bin/bash
include .env

# Lint is to fix bad code practices
lint: ## Run linters only.
	@echo "\033[2m→ Running linters...\033[0m"
	@golangci-lint run --config .golangci.yml --fix

# run the project
run: ## Run the application.
	@echo "\033[2m→ Starting the project...\033[0m"
	@docker start mysqldb_fiber_site
	@go run .

# reset the database
resetdb:
	@echo "\033[2m→ Resetting the database...\033[0m"
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -h $(DB_HOST) $(DB_NAME) < ./scripts/reset.sql

# build the project
build:
	@echo "\033[2m→ Building the project...\033[0m"
	@go build -o ./bin/fiber_site .
	
# serve the project from the binary
serve:
	@echo "\033[2m→ Starting the server...\033[0m"
	@mkdir -p ./log
	@./bin/fiber_site >./log/`date +%F`.log 2>&1 &

