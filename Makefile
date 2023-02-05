test:
	go test ./...

lint:
	golangci-lint run

run:
	go run cmd/app/main.go

down:
	docker-compose down

up:
	docker-compose up -d

logs:
	docker-compose logs 
