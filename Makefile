#!/bin/bash
include .env

# Lint is to fix bad code practices
lint: ## Run linters only.
	@echo "\033[2m→ Running linters...\033[0m"
	@golangci-lint run --config .golangci.yml --fix
	
# go formater
fmt:
	@echo "\033[2m→ Running gofmt...\033[0m"
	@gofmt -l -s -w .

# run the project
run: ## Run the application.
	@echo "\033[2m→ Starting the project...\033[0m"
	@docker start mysqldb_fiber_site
	@go run .

enterdb:
	@echo "\033[2m→ Entering the DB of the project...\033[0m"
	@docker exec -it mysqldb_fiber_site mysql -u root -ppassword

# reset the database
resetdb:
	@echo "\033[2m→ Resetting the database...\033[0m"
	@mysql -u $(DB_USER) -p$(DB_PASSWORD) -h $(DB_HOST) < ./scripts/reset.sql

# build the project
build:
	@echo "\033[2m→ Building the project...\033[0m"
	@go build -o ./bin/fiber_site .
	
# serve the project from the binary
serve:
	@echo "\033[2m→ Starting the server...\033[0m"
	@mkdir -p ./log
	@ ./bin/fiber_site --port 80 >./log/`date +%F`.log 2>./log/`date +%F`-err.log  &

# serve the project from the binary
serveTLS:
	@echo "\033[2m→ Starting the server...\033[0m"
	@mkdir -p ./log
	@ ./bin/fiber_site --port 443 -TLS >./log/`date +%F`-TLS.log 2>./log/`date +%F`-errTLS.log  &

# dump the sql
dump:
	@echo "\033[2m→ Dumping the database...\033[0m"
	@docker exec -it mysqldb_fiber_site mysqldump -u root -ppassword gofiber_website > dump.sql

