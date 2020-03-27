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
## Create migration. Usage: 'make migration-create name=some-name'
migration-create: ; $(info creating migration...) @ ## Create migration
	@ret=0; if [ -z $(name) ]; then \
		ret=1; \
		echo "Migration name not specified"; \
	else \
		go run cmd/migration/main.go create $(name); \
	fi; exit $$ret

.PHONY: migration-run
## Run migrations. Usage: 'make migration-run dir=(up|down) count=<number>[opt]'
migration-run: ; $(info running migrations...) @ ## Create migration
	@ret=0; if [ -z $(dir) ]; then \
		ret=1; \
		echo "Migration direction not specified (up|down)"; \
	elif [ "$(dir)" != "up" ] && [ "$(dir)" != "down" ]; then \
		ret=1; \
		echo "Invalid direction provided: '$(dir)'"; \
	else \
		if ! [ -z $(count) ]; then \
			if [ "`echo $(count) | egrep ^[1-9][0-9]*$$`" = "" ]; then \
				ret=1; \
				echo "Invalid count provided: '$(count)'"; \
			else \
				go run cmd/migration/main.go $(dir) $(count); \
			fi; \
		else \
			go run cmd/migration/main.go $(dir); \
		fi; \
	fi; exit $(ret)

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