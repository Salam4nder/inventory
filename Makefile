test:
	go test ./...

lint:
	golangci-lint run

run:
	go run cmd/app/main.go
