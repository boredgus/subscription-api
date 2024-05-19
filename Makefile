compose_file:=docker/docker-compose.yaml

start:
	docker compose -f $(compose_file)  --env-file docker/default.env up

generate-mocks:
	mockery --config=config/.mockery.yaml

test:
	go test ./... -coverprofile="test-coverage.txt" -covermode count
	go tool cover -func="test-coverage.txt"

test-coverage:
	go tool cover -html="test-coverage.txt"

generate-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/grpc/currency_service/currencyService.proto \
		pkg/grpc/dispatch_service/dispatchService.proto