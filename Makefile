## go-build :
go-build:
	go build -o bin/main main.go
go-run:
	go run cmd/app/main.go
docker-run:
	docker-compose up --build -d
go-test:
	go test