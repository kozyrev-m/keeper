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

.DEFAULT_GOAL= hello