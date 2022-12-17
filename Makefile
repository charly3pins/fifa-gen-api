migration-run: export POSTGRESQL_URL=postgresql://fifa_gen_dev:fifa_gen_dev@localhost:5431/fifa_gen_dev_db?sslmode=disable

.PHONY: run
## Run service. Usage: 'make run'
run: ; $(info running code…) @
	go run ./cmd/server/main.go

.PHONY: up
## start DB in Docker. Usage: 'make up'
up: ; $(info starting db…) @
	docker-compose -f ./docker-compose.yaml up -d; \
	sleep 5s; \
	go run cmd/migration/main.go up;

## stop DB in Docker. Usage: 'make down'
down: ; $(info starting db…) @
	docker-compose -f ./docker-compose.test.yml down;

.PHONY: migration-create
## Creates a new migration usage: `migration-create name=<migration name>`
migration-create:
	@migrate create -dir ./cmd/migration/sqls -ext sql $(name)

.PHONY: migration-run
## Runs migrations: `migration-run dir=[up,down] (optional count=[number of migrations])`
migration-run:
	$(info Running migrations...)
	@migrate -database ${POSTGRESQL_URL} -path ./cmd/migration/sqls $(dir) $(count)

# -- help

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=20
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)