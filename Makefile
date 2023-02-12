create-client:
	go build -v -o keeper-client ./cmd/client/main.go

test:
	@sudo docker-compose --file docker/test/docker-compose.yaml up -d
	@go test ./...
	@sudo docker-compose --file docker/test/docker-compose.yaml down

test-coverage:
	@sudo docker-compose --file docker/test/docker-compose.yaml up -d
	@echo "Package test coverage:"
	@go test -coverpkg=./internal/... -coverprofile=coverage.out ./...
	@echo "\n\n"
	@echo "Separate files test coverage:"
	@go tool cover -func coverage.out
	@sudo docker-compose --file docker/test/docker-compose.yaml down
	@timeout 5 echo
	@rm coverage.out

hello:
	@echo "Use 'make' with a specific commands"
	@echo " "
	@echo "Commands: "
	@echo " create-client		create client to connect to keeper"
	@echo " test			testing"
	@echo " test-coverage		another testing statistic"

migrations-up:
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" up
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" up

migrations-down:
	@goose -dir=./migrations postgres "host=localhost port=5433 dbname=keeper_dev user=admin password=12345 sslmode=disable" down
	@goose -dir=./migrations postgres "host=localhost port=5434 dbname=keeper_test user=admin password=12345 sslmode=disable" down

.DEFAULT_GOAL= hello