GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Build:
build-client: ## build client and put binary to bin/
	go build -v -o ./keeper-client ./cmd/client/main.go

## Docker:
docker-keeper-up: ## create and run containers with keeper and database
	docker-compose --file build/docker-compose.yml up -d

docker-keeper-down: ## stop and remove containers with keeper and database
	docker-compose --file build/docker-compose.yml down

## Test:
test: ## run tests
	@sudo docker-compose --file docker/test/docker-compose.yaml up -d
	@go test ./...
	@sudo docker-compose --file docker/test/docker-compose.yaml down

test-coverage: ## run test and show coverage
	@sudo docker-compose --file docker/test/docker-compose.yaml up -d
	@echo "Package test coverage:"
	@go test -coverpkg=./internal/... -coverprofile=coverage.out ./...
	@echo "\n\n"
	@echo "Separate files test coverage:"
	@go tool cover -func coverage.out
	@sudo docker-compose --file docker/test/docker-compose.yaml down
	@timeout 5 echo
	@rm coverage.out

migrations-up: 
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" up
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" up

migrations-down:
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" down
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" down

## Help:
help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

.DEFAULT_GOAL= hello