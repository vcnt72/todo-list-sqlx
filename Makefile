go-build:
	go build -o bin/main main.go
go-run:
	go run cmd/app/main.go
docker-run:
	docker-compose up --build -d
go-test:
	go test
migrate-up: 
	migrate -path db/migrations -database "postgres://postgres:mateup123@localhost:5432/todo_list?sslmode=disable" -verbose up
migrate-down: #Need migrate CLI https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
	migrate -path db/migrations -database "postgres://postgres:mateup123@localhost:5432/todo_list?sslmode=disable" -verbose down