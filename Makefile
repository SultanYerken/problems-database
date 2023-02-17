build:
	docker-compose up --build problems-database

run:
	docker-compose up problems-database

migrate:
	migrate -path db/migrations -database postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable up
	
