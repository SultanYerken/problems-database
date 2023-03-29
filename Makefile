build:
	docker-compose up --build problems-database

run:
	docker-compose up problems-database

migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path db/migrations -database postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable up
	
