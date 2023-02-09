test:
	go test ./...

test-coverage:
	@echo "Package test coverage:"
	@go test -coverpkg=./internal/... -coverprofile=coverage.out ./...
	@echo "\n\n"
	@echo "Separate files test coverage:"
	@go tool cover -func coverage.out

hello:
	@echo "Use 'make' with a specific command:"
	@echo "1. test"
	@echo "2. test-coverage"

migrations-up:
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" up
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" up

migrations-down:
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" down
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" down

.DEFAULT_GOAL= hello